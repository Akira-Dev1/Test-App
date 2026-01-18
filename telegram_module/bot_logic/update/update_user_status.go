package update

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"bot_logic/storage"
)

type TokenResponse struct {
	Status       string `json:"status"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}

func UpdateUserStatus() {
	log.Println("=== START UpdateUserStatus ===")
	defer log.Println("=== END UpdateUserStatus ===")
	
	ctx := context.Background()
	
	// Проверяем соединение с Redis
	log.Println("Pinging Redis...")
	if err := storage.AuthRedis.Ping(ctx).Err(); err != nil {
		log.Printf("Redis connection error: %v", err)
		return
	}
	log.Println("Redis connection OK")
	
	// Логируем перед получением данных
	log.Println("Getting IDs from Redis...")
	ids, err := storage.AuthRedis.HGetAll(ctx, "ids").Result()
	if err != nil {
		log.Printf("Error getting IDs from Redis: %v", err)
		return
	}
	
	log.Printf("Found %d IDs to process", len(ids))
	
	if len(ids) == 0 {
		log.Println("No IDs to process, exiting")
		return
	}

	// Выводим все ID для отладки
	log.Printf("IDs found: %v", ids)
	
	for id := range ids {
		log.Printf("=== Processing user ID: %s ===", id)
		
		entry_token, err := storage.AuthRedis.HGet(ctx, id, "entry_token").Result()
		if err != nil {
			log.Printf("Error getting entry_token for %s: %v", id, err)
			continue
		}

		log.Printf("Got entry_token: %s", entry_token)
		
		url := "http://auth:8081/login/check?entry_token=" + entry_token
		log.Printf("Making request to auth service: %s", url)

		// Создаем HTTP клиент с таймаутом
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		
		startTime := time.Now()
		resp, err := client.Get(url)
		if err != nil {
			log.Printf("Error making request for %s: %v", id, err)
			continue
		}
		defer resp.Body.Close()

		log.Printf("Request completed in %v, status: %s", time.Since(startTime), resp.Status)

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response for %s: %v", id, err)
			continue
		}

		log.Printf("Raw response from auth (%d bytes): %s", len(bodyBytes), string(bodyBytes))

		var tokenResp TokenResponse
		if err := json.Unmarshal(bodyBytes, &tokenResp); err != nil {
			log.Printf("Error parsing JSON for %s: %v", id, err)
			log.Printf("Response that failed to parse: %s", string(bodyBytes))
			continue
		}

		log.Printf("Parsed status: %s, UserID: %s", tokenResp.Status, tokenResp.UserID)

		switch tokenResp.Status {
		case "approved":
			log.Printf("Setting authorized for user %s", id)
			err := storage.AuthRedis.HSet(ctx, id,
				"access_token", tokenResp.AccessToken,
				"refresh_token", tokenResp.RefreshToken,
				"user_id", tokenResp.UserID,
				"status", "authorized",
			).Err()
			if err != nil {
				log.Printf("Error setting authorized for %s: %v", id, err)
			} else {
				log.Printf("Successfully authorized user %s", id)
			}
			storage.AuthRedis.HDel(ctx, "ids", id)

		case "denied":
			log.Printf("Setting denied for user %s", id)
			storage.AuthRedis.HSet(ctx, id, "status", "unknown")
			storage.AuthRedis.HDel(ctx, "ids", id)

		case "pending":
			log.Printf("User %s is still pending", id)
			// Ничего не делаем, оставляем в ids

		default:
			log.Printf("Unknown status '%s' for user %s", tokenResp.Status, id)
		}
		
		log.Printf("=== Finished processing user ID: %s ===", id)
	}
}