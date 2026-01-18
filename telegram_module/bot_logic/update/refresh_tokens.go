package update

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"io"

	"bot_logic/storage"
)

type RefreshResponseStruct struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func RefreshUsersTokens() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for range ticker.C {
			ctx := context.Background()

			iter := storage.AuthRedis.Scan(ctx, 0, "*", 0).Iterator()

			for iter.Next(ctx) {
				chat_id := iter.Val()
				if chat_id == "ids" {
					continue
				}

				url := "http://auth:8081/refresh"
				id, _ := strconv.Atoi(chat_id)
				userData, err := storage.AuthRedis.HGetAll(ctx, chat_id).Result()
				if err != nil {
					log.Printf("[ERROR] Failed to HGETALL for chatID %s: %v", chat_id, err)
					continue
				}

				status := userData["status"]

				if status != "authorized" {
					log.Printf("[DEBUG] Skipping chatID %s: status is '%s' (need 'authorized')", chat_id, status)
					continue
				}

				refresh_token, err := storage.GetRefreshToken(id)
				if err != nil {
					log.Printf("Status update failed: %v", err)
					continue
				}

				data := map[string]string{"refresh_token": refresh_token}
				jsonData, err := json.Marshal(data)
				if err != nil {
					log.Printf("Ошибка создания JSON: %v\n", err)
					continue
				}

				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
				if err != nil {
					log.Printf("Ошибка создания запроса: %v\n", err)
					continue
				}
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					log.Printf("Ошибка запроса: %v\n", err)
					continue
				}

				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Error reading response body: %v", err)
					continue
				}

				log.Print(string(bodyBytes))

				var refreshResp RefreshResponseStruct
				if err := json.Unmarshal(bodyBytes, &refreshResp); err != nil {
					log.Printf("Error decoding JSON: %v", err)
					continue
				}

				err = storage.AuthRedis.HSet(ctx, chat_id, "access_token", refreshResp.AccessToken, "refresh_token", refreshResp.RefreshToken).Err()
				if err != nil {
					log.Printf("Error: %v", err)
					continue
				}
			}
		}
	}()
}
