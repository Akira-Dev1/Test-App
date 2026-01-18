package main_module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Получить данные о пользователе (курсы, оценки, тесты)
func GetUserData(token, targetID string, courses, tests, grades bool) (*UserProfileData, error) {
	url := fmt.Sprintf("%s/users/%s/data?c=%t&t=%t&g=%t", baseURL, targetID, courses, tests, grades)
	req, _ := http.NewRequest("GET", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data UserProfileData
	json.NewDecoder(resp.Body).Decode(&data)
	return &data, nil
}

// Получить список пользователей
func GetUserList(token string) ([]UserInfo, error) {
	req, _ := http.NewRequest("GET", baseURL+"/users", nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var list []UserInfo
	json.NewDecoder(resp.Body).Decode(&list)
	return list, nil
}


// Получить ФИО пользователя
func GetUserInfo(token, targetID string) (map[string]any, error) {
	url := fmt.Sprintf("%s/users/%s", baseURL, targetID)
	req, _ := http.NewRequest("GET", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: status %d", resp.StatusCode)
	}

	var res map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// Изменить ФИО пользователя
func UpdateUserFullName(token, targetID, newName string) error {
	payload := map[string]string{"full_name": newName}
	body, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s/users/%s/name", baseURL, targetID)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Получить роли пользователя
func GetUserRoles(token, targetID string) ([]string, error) {
	url := fmt.Sprintf("%s/users/%s/roles", baseURL, targetID)
	req, _ := http.NewRequest("GET", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res UserRolesResponse
	json.NewDecoder(resp.Body).Decode(&res)
	return res.Roles, nil
}

// Изменить роли пользователя
func UpdateUserRoles(token, targetID string, roles []string) error {
	payload := map[string][]string{"roles": roles}
	body, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s/users/%s/roles", baseURL, targetID)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update roles: status %d", resp.StatusCode)
	}
	return nil
}

// Получить статус пользователя (заблокирован/разблокирован)
func GetUserBlockStatus(token, targetID string) (bool, error) {
	url := fmt.Sprintf("%s/users/%s/block-status", baseURL, targetID)
	req, _ := http.NewRequest("GET", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to get block status: status %d", resp.StatusCode)
	}

	var res struct {
		IsBlocked bool `json:"is_blocked"`
	}
	json.NewDecoder(resp.Body).Decode(&res)
	return res.IsBlocked, nil
}

// Заблокировать/Разблокировать пользователя
func SetUserBlockStatus(token, targetID string, isBlocked bool) error {
	payload := map[string]bool{"is_blocked": isBlocked}
	body, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s/users/%s/block", baseURL, targetID)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to set block status: status %d", resp.StatusCode)
	}
	return nil
}
