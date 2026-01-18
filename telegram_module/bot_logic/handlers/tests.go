package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"bot_logic/main_module"
)

// Получение тестов
func GetTestsHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	courseID, _ := strconv.Atoi(r.URL.Query().Get("course_id"))

	tests, err := main_module.GetCourseTests(token, courseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tests)
}

// Создание теста
func CreateTestHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	courseID, _ := strconv.Atoi(r.URL.Query().Get("course_id"))

	var body struct {
		Title string `json:"title"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	id, err := main_module.CreateTest(token, courseID, body.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// Удаление теста
func DeleteTestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token, _ := getToken(r)

	testIDStr := r.URL.Query().Get("test_id")
	testID, err := strconv.Atoi(testIDStr)

	if err != nil {
		http.Error(w, "Invalid test_id", http.StatusBadRequest)
		return
	}

	err = main_module.DeleteTest(token, testID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Получение статуса теста
func GetTestStatusHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	courseID, _ := strconv.Atoi(r.URL.Query().Get("course_id"))
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))

	isActive, err := main_module.GetTestStatus(token, courseID, testID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"is_active": isActive})
}

// Добавить вопрос в тест
func AddQuestionToTestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, _ := getToken(r)
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))
	questionID, _ := strconv.Atoi(r.URL.Query().Get("question_id"))

	err := main_module.AddQuestionToTest(token, testID, questionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Удалить вопрос из теста
func RemoveQuestionFromTestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, _ := getToken(r)
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))
	questionID, _ := strconv.Atoi(r.URL.Query().Get("question_id"))

	err := main_module.RemoveQuestionFromTest(token, testID, questionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Активация/Деактивация теста
func ToggleTestActivationHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	courseID, _ := strconv.Atoi(r.URL.Query().Get("course_id"))
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))

	var body struct {
		IsActive bool `json:"is_active"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := main_module.UpdateTestActivation(token, courseID, testID, body.IsActive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Изменение порядка вопросов
func ReorderQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	testID, _ := strconv.Atoi(r.URL.Query().Get("test_id"))

	var body struct {
		QuestionIDs []int `json:"question_ids"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := main_module.ReorderQuestions(token, testID, body.QuestionIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}