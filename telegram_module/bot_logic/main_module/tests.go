package main_module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Получить тесты по айди курса
func GetCourseTests(token string, courseID int) ([]Test, error) {
	url := fmt.Sprintf("%s/courses/%d/tests", baseURL, courseID)
	req, _ := http.NewRequest("GET", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get tests: %d", resp.StatusCode)
	}

	var data TestsListResponse
	json.NewDecoder(resp.Body).Decode(&data)
	return data.Tests, nil
}

// Создать тест
func CreateTest(token string, courseID int, title string) (int, error) {
	payload := map[string]string{"title": title}
	body, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s/courses/%d/tests", baseURL, courseID)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("failed to create test: %d", resp.StatusCode)
	}

	var res struct{ ID int `json:"id"` }
	json.NewDecoder(resp.Body).Decode(&res)
	return res.ID, nil
}

// Удалить тест
func DeleteTest(token string, testID int) error {
	url := fmt.Sprintf("%s/tests/%d", baseURL, testID)
	req, _ := http.NewRequest("DELETE", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete test: %d", resp.StatusCode)
	}
	return nil
}

// Получить статус теста
func GetTestStatus(token string, courseID, testID int) (bool, error) {
	url := fmt.Sprintf("%s/courses/%d/tests/%d/status", baseURL, courseID, testID)
	req, _ := http.NewRequest("GET", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to get test status: %d", resp.StatusCode)
	}

	var res TestStatusResponse
	json.NewDecoder(resp.Body).Decode(&res)
	return res.IsActive, nil
}

// Активировать/Деактивировать тест
func UpdateTestActivation(token string, courseID, testID int, isActive bool) error {
	payload := map[string]bool{"is_active": isActive}
	body, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s/courses/%d/tests/%d/activation", baseURL, courseID, testID)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to update test status: %d", resp.StatusCode)
	}
	return nil
}

// Добавление вопроса в тест
func AddQuestionToTest(token string, testID, questionID int) error {
	url := fmt.Sprintf("%s/tests/%d/questions/%d", baseURL, testID, questionID)
	req, _ := http.NewRequest("POST", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to add question: %d", resp.StatusCode)
	}
	return nil
}

// Удаление вопроса с теста
func RemoveQuestionFromTest(token string, testID, questionID int) error {
	url := fmt.Sprintf("%s/tests/%d/questions/%d", baseURL, testID, questionID)
	req, _ := http.NewRequest("DELETE", url, nil)
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to remove question: %d", resp.StatusCode)
	}
	return nil
}

// Изменение порядка вопросов в тесте
func ReorderQuestions(token string, testID int, questionIDs []int) error {
	payload := map[string][]int{"question_ids": questionIDs}
	body, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s/tests/%d/questions/reorder", baseURL, testID)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	setAuthHeader(req, token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to reorder questions: %d", resp.StatusCode)
	}
	return nil
}