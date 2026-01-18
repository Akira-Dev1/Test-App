package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Notification struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type AllNotificationsResponse map[string][]Notification

func GetUserNotifications() {
	ticker := time.NewTicker(10 * time.Second)

	client := &http.Client{
		Timeout: 8 * time.Second,
	}

	go func() {
		for range ticker.C {
			url := "http://tg_nginx/notifications"

			resp, err := client.Get(url)
			if err != nil {
				log.Printf("Failed to fetch notifications from Bot Logic: %v", err)
				continue
			}

			var all_notifications AllNotificationsResponse
			err = json.NewDecoder(resp.Body).Decode(&all_notifications)

			resp.Body.Close()

			if err != nil {
				log.Printf("Decoding failed: %v", err)
				continue
			}

			if len(all_notifications) == 0 {
				continue
			}

			for chatID, notifications := range all_notifications {
				for _, n := range notifications {
					botToken := os.Getenv("BOT_TOKEN")
					if botToken == "" {
						log.Fatal("BOT_TOKEN not set")
					}
					url := "https://api.telegram.org" + botToken + "/sendMessage"

					fullText := fmt.Sprintf("<b>%s</b>\n\n%s", n.Title, n.Message)
					payload := map[string]interface{}{
						"chat_id":    chatID,
						"text":       fullText,
						"parse_mode": "HTML",
					}

					body, _ := json.Marshal(payload)

					resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
					if err != nil {
						log.Printf("Error sending to TG API: %v", err)
						return
					}
					defer resp.Body.Close()

					log.Printf("Notifications check sent to %s, Response: %d", url, resp.StatusCode)

					if resp.StatusCode != http.StatusOK {
						log.Printf("TG API returned error: %d", resp.StatusCode)
					}

					log.Printf("Sending to Chat %s: [%s] %s", chatID, n.Title, n.Message)
				}
			}
		}
	}()
}
