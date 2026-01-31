package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	authTransport "main_module/internal/auth/transport"
	testDomain "main_module/internal/test/domain"
)


// Удалить вопрос из теста
func (h *Handler) RemoveQuestionFromTest(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	questionIDString := r.PathValue("questionID")
	testIDString := r.PathValue("testID")
	if questionIDString == "" || testIDString == "" {
		http.Error(w, "missing questionID or testID", http.StatusBadRequest)
	}
	questionID, _ := strconv.Atoi(questionIDString)
	testID, _ := strconv.Atoi(testIDString)

	err := h.TestService.RemoveQuestionFromTest(&user, testID, questionID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}
// Добавить вопрос в тест
func (h *Handler) AddQuestionToTest(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	questionIDString := r.PathValue("questionID")
	testIDString := r.PathValue("testID")
	if questionIDString == "" || testIDString == "" {
		http.Error(w, "missing questionID or testID", http.StatusBadRequest)
	}
	questionID, _ := strconv.Atoi(questionIDString)
	testID, _ := strconv.Atoi(testIDString)

	err := h.TestService.AddQuestionToTest(&user, testID, questionID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}
// Изменить порядок следования вопросов в тесте
func (h *Handler) ChangeOrderOfQuestions(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	testIDString := r.PathValue("testID")
	if testIDString == "" {
		http.Error(w, "missing questionID or testID", http.StatusBadRequest)
	}
	testID, _ := strconv.Atoi(testIDString)

	var data map[string]any

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}
	var questionData testDomain.TestQuestionIDS
	questionIDS, hasQuestionIDS := data["question_ids"].([]int)

	if !hasQuestionIDS {
		http.Error(w, "not enought question_ids field", 400)
		return
	}
	questionData.IDS = questionIDS

	err := h.TestService.ChangeOrderOfQuestions(&user, testID, questionData)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}
// Получить список вопросов в тесте
func (h *Handler) GetTestQuestions(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	testIDString := r.PathValue("testID")
	if testIDString == "" {
		http.Error(w, "missing questionID or testID", http.StatusBadRequest)
	}
	testID, _ := strconv.Atoi(testIDString)

	questionIDS, err := h.TestService.GetTestQuestions(&user, testID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questionIDS)
}