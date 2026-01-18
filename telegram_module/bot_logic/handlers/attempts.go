package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"bot_logic/main_module"
)

// Получить прошедших юзеров
func GetPassedUsersHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))
	
	data, err := main_module.GetPassedUsers(token, testID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// Получить оценку
func GetScoresHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))

	data, err := main_module.GetTestScores(token, testID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// Получить ответы
func GetAnswersHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))

	data, err := main_module.GetTestAnswers(token, testID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// Начать попытку
func StartAttemptHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))

	id, err := main_module.StartAttempt(token, testID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"attempt_id": id})
}

// Обновить ответ
func SubmitAnswerHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	attemptID, _ := strconv.Atoi(r.URL.Query().Get("attempt_id"))

	var body struct {
		QuestionID  int `json:"question_id"`
		AnswerIndex int `json:"answer_index"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	err := main_module.SubmitAnswer(token, attemptID, body.QuestionID, body.AnswerIndex)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(204)
}

// Удалить ответ
func DeleteAnswerHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	attemptID, _ := strconv.Atoi(r.URL.Query().Get("attempt_id"))
	questionID, _ := strconv.Atoi(r.URL.Query().Get("question_id"))

	err := main_module.DeleteAttemptAnswer(token, attemptID, questionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Обновить ответ на конкретный вопрос
func UpdateAnswerHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	attemptID, _ := strconv.Atoi(r.URL.Query().Get("attempt_id"))
	questionID, _ := strconv.Atoi(r.URL.Query().Get("question_id"))

	var body struct {
		AnswerIndex int `json:"answer_index"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	err := main_module.UpdateAttemptAnswer(token, attemptID, questionID, body.AnswerIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Завершить попытку
func CompleteAttemptHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	attemptID, _ := strconv.Atoi(r.URL.Query().Get("attempt_id"))

	err := main_module.CompleteAttempt(token, attemptID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Посмотреть попытку
func GetUserAttemptHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))
	targetUserID := r.URL.Query().Get("target_user_id")

	data, err := main_module.GetAttemptData(token, testID, targetUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// Посмотреть ответы пользователя
func GetUserAttemptAnswersHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))
	targetUserID := r.URL.Query().Get("target_user_id")

	data, err := main_module.GetAttemptAnswers(token, testID, targetUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(data)
}