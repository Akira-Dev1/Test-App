package auth

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"io"
	"fmt"
)


func YandexLogin(chat_id int, entry_token string) (string, error) {
	url := "http://auth:8081/login?type=yandex"

	body := map[string]string{"entry_token": entry_token}
	jsonData, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Err request: %v\n", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}	
    resp, err := client.Do(req)
	if err != nil {
		log.Printf("Err response: %v\n", err)
		return "", err
	}
    defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response: %v", err)
    }
	fmt.Println(string(bodyBytes))

    return string(bodyBytes), nil
}


