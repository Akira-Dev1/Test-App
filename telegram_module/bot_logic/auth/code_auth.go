package auth

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type CodeResponse struct {
	Code string `json:"code"`
}

func CodeAuth(chat_id int, entry_token string) (string, error) {

	url := "http://auth:8081/login?type=code"

	body := map[string]string{"entry_token": entry_token}
	jsonData, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Err request: %v\n", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Err response: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	var coderesp CodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&coderesp); err != nil {
		log.Printf("Err request: %v\n", err)
		return "", err
	}

	return coderesp.Code, nil
}
