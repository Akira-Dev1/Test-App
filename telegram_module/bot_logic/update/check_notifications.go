package update

import (
	"context"
	"bot_logic/storage"
	"log"
)

func CollectNotificationsFromAllUsers() (map[string][]Notification, error) {
	ctx := context.Background()

	log.Println("[INFO] === Starting Global Notification Collection ===")

	user_notifications := make(map[string][]Notification)

	iter := storage.AuthRedis.Scan(ctx, 0, "*", 0).Iterator()

	for iter.Next(ctx) {
		chatID := iter.Val()

		userData, err := storage.AuthRedis.HGetAll(ctx, chatID).Result()
		if err != nil {
			log.Printf("[ERROR] Failed to HGETALL for chatID %s: %v", chatID, err)
			continue
		}

		status := userData["status"]
		token := userData["access_token"]

		if status != "authorized" {
			log.Printf("[DEBUG] Skipping chatID %s: status is '%s' (need 'authorized')", chatID, status)
			continue
		}

		if token == "" {
			log.Printf("[WARN] Skipping chatID %s: status is 'authorized' but access_token is EMPTY", chatID)
			continue
		}

		log.Printf("[DEBUG] Processing authorized user: chatID %s (userID %s)", chatID, userData["user_id"])

		notifications, err := fetchUserNotifications(token)
		if err != nil {
			log.Printf("[ERROR] Failed to fetch notifications for chatID %s: %v", chatID, err)
			continue
		}

		if len(notifications) == 0 {
			log.Printf("[DEBUG] No new notifications for chatID %s", chatID)
			continue
		}

		user_notifications[chatID] = notifications
		log.Printf("[SUCCESS] Found %d notifications for chatID %s", len(notifications), chatID)

		var ids []int
		for _, n := range notifications {
			ids = append(ids, n.ID)
		}
		confirmNotifications(token, ids)
	}
	if err := iter.Err(); err != nil {
		log.Printf("[CRITICAL] Redis SCAN iterator error: %v", err)
		return nil, err
	}
	return user_notifications, nil
}