package infrastructure

import (
	"errors"
	"time"
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	userDomain 	"main_module/internal/user/domain"
)


// Получить список пользователей
func  (r *PostgresRepo) GetUsersIDS() (userDomain.UsersIDS, error) {
    client := &http.Client{
        Timeout: 3 * time.Second,
    }

    req, err := http.NewRequest(
        "GET",
        "http://auth:8081/userservice/get_user_list",
        nil,
    )
    if err != nil {
        return userDomain.UsersIDS{}, errors.New("failed to create request")
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return userDomain.UsersIDS{}, errors.New("request failed")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return userDomain.UsersIDS{}, fmt.Errorf("auth service error: status %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return userDomain.UsersIDS{}, fmt.Errorf("failed to read response body: %w", err)
    }

    var users userDomain.UsersIDS
    if err := json.Unmarshal(body, &users); err != nil {
        return userDomain.UsersIDS{}, fmt.Errorf("failed to unmarshal JSON: %w, body: %s", err, string(body))
    }

    return users, nil
}

// Получить информацию о пользователе (Список ролей)
func  (r *PostgresRepo) GetRoles(targetUserID string) (userDomain.UserData, error) {
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
        "http://auth:8081/userservice/get_user_roles",
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

    var roles []string
    if err := json.Unmarshal(body, &roles); err != nil {
        return userDomain.UserData{}, fmt.Errorf("failed to unmarshal JSON: %w, body: %s", err, string(body))
    }
	userData := userDomain.UserData{
		Roles: roles,
	}

    return userData, nil
}

// Изменить информацию о пользователе (Список ролей)
func  (r *PostgresRepo) UpdateRoles(targetUserID string, roles []string) error {
    requestBody := map[string]interface{} {
        "user_id": targetUserID,
		"roles": roles,
    }
    
    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        return errors.New("failed to marshal JSON")
    }

    client := &http.Client{
        Timeout: 3 * time.Second,
    }

    req, err := http.NewRequest(
        "PATCH",
        "http://auth:8081/userservice/update_user_roles",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        return errors.New("failed to create request")
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return errors.New("request failed")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return fmt.Errorf("auth service error: status %d", resp.StatusCode)
    }

    return nil
}

// Получить информацию о пользователе (Заблокирован ли пользователь)
func  (r *PostgresRepo) GetStatus(targetUserID string) (userDomain.UserData, error) {
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
        "http://auth:8081/userservice/get_user_block_status",
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

    var data map[string]bool
    if err := json.Unmarshal(body, &data); err != nil {
        return userDomain.UserData{}, fmt.Errorf("failed to unmarshal JSON: %w, body: %s", err, string(body))
    }
	isBlocked := data["Is_blocked"]
	userData := userDomain.UserData{
		IsBlocked: &isBlocked,
	}

    return userData, nil
}

// Изменить информацию о пользователе (Заблокировать/Разблокировать пользователя)
func  (r *PostgresRepo) UpdateStatus(targetUserID string, status bool) error {
    requestBody := map[string]interface{} {
        "user_id": targetUserID,
		"is_blocked": status,
    }
    
    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        return errors.New("failed to marshal JSON")
    }

    client := &http.Client{
        Timeout: 3 * time.Second,
    }

    req, err := http.NewRequest(
        "POST",
        "http://auth:8081/userservice/set_block_status",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        return errors.New("failed to create request")
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return errors.New("request failed")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return fmt.Errorf("auth service error: status %d", resp.StatusCode)
    }

    return nil
}