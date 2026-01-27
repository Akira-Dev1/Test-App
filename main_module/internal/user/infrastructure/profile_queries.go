package infrastructure

import (
	"errors"
	"net/http"
	"encoding/json"
	"bytes"
	"time"
	"fmt"
	"io"

	userDomain 	"main_module/internal/user/domain"
)

// Получить информацию о пользователе (ФИО)
func  (r *PostgresRepo) GetName(targetUserID string) (userDomain.UserData, error) {
    requestBody := map[string]string{
        "user_id": targetUserID,
    }
    
    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        return userDomain.UserData{}, errors.New("failed to marshal JSON")
    }

    client := &http.Client{
        Timeout: 3 * time.Second,
    }

    req, err := http.NewRequest(
        "GET",
        "http://auth:8081/userservice/get_user_info",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        return userDomain.UserData{}, errors.New("failed to create request")
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return userDomain.UserData{}, errors.New("request failed")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return userDomain.UserData{}, fmt.Errorf("auth service error: status %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return userDomain.UserData{}, fmt.Errorf("failed to read response body: %w", err)
    }

    var data map[string]string
    if err := json.Unmarshal(body, &data); err != nil {
        return userDomain.UserData{}, fmt.Errorf("failed to unmarshal JSON: %w, body: %s", err, string(body))
    }
	name := data["Name"]
	userData := userDomain.UserData{
		Name: &name,
	}

    return userData, nil
}

// Изменить информацию о пользователе (ФИО)
func  (r *PostgresRepo) UpdateName(targetUserID string, name string) error {
	requestBody := map[string]string{
        "user_id":  targetUserID,
        "new_name": name,
    }

    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        return fmt.Errorf("failed to marshal JSON: %w", err)
    }

    client := &http.Client{
        Timeout: 3 * time.Second,
    }

    req, err := http.NewRequest(
        "PATCH",
        "http://auth:8081/userservice/update_full_name",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        return fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("request failed: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == 200 {
        return nil
    } else {
        return fmt.Errorf("error from auth-service, status: %d", resp.StatusCode)
    }
}