package update

import (
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)

type Notification struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

type MainModuleResponse struct {
	Notifications []Notification `json:"notifications"`
}

// Функция получения уведомлений
func fetchUserNotifications(token string) ([]Notification, error) {
	url := "http://main_module:18080/notification"

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
		log.Printf("[ERROR] Failed to create request: %v", err)
        return nil, err
    }

	req.Header.Set("Authorization", "Bearer "+token)

	log.Printf("[DEBUG] Fetching notifications from: %s", url)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[ERROR] Network error reaching main_module: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("[DEBUG] Main module responded with status: %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("main module returned non-200 status: %d", resp.StatusCode)
	}

	var res MainModuleResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Printf("[ERROR] Failed to decode JSON response: %v", err)
		return nil, err
	}

	log.Printf("[INFO] Successfully received %d notifications", len(res.Notifications))

	return res.Notifications, nil
}


// Функция подтверждения получения уведомлений
func confirmNotifications(token string, ids []int) {
	url := "http://main_app_container:8080/notification/confirm-tg"

	body, _ := json.Marshal(map[string][]int{"ids": ids})
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err == nil {
		resp.Body.Close()
	}
}