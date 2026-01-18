package main_module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Создать вопрос
func CreateQuestion(token string, q Question) (int, error) {
	body, _ := json.Marshal(q)
	req, _ := http.NewRequest("POST", baseURL+"/questions", bytes.NewBuffer(body))
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("failed to create question: %d", resp.StatusCode)
	}

	var res struct{ ID int `json:"id"` }
	json.NewDecoder(resp.Body).Decode(&res)
	return res.ID, nil
}

// Удалить вопрос
func DeleteQuestion(token string, id int) error {
	url := fmt.Sprintf("%s/questions/%d", baseURL, id)
	req, _ := http.NewRequest("DELETE", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete question (maybe used in tests): %d", resp.StatusCode)
	}
	return nil
}

// Изменить вопрос
func UpdateQuestion(token string, id int, q Question) (int, error) {
	body, _ := json.Marshal(q)
	url := fmt.Sprintf("%s/questions/%d", baseURL, id)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to update: %d", resp.StatusCode)
	}

	var res struct{ NewVersion int `json:"new_version"` }
	json.NewDecoder(resp.Body).Decode(&res)
	return res.NewVersion, nil
}

// Получить информацию о вопросе
func GetQuestionDetail(token string, id, version int) (*Question, error) {
	url := fmt.Sprintf("%s/questions/%d/%d", baseURL, id, version)
	req, _ := http.NewRequest("GET", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("question not found or denied: %d", resp.StatusCode)
	}

	var q Question
	json.NewDecoder(resp.Body).Decode(&q)
	return &q, nil
}

// Получить вопросы
func GetMyQuestions(token string) ([]Question, error) {
	req, _ := http.NewRequest("GET", baseURL+"/questions", nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data QuestionListResponse
	json.NewDecoder(resp.Body).Decode(&data)
	return data.Questions, nil
}
