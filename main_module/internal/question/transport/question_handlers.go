package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	authTransport "main_module/internal/auth/transport"
)

// Получить список вопросов
func (h *Handler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}


	questions, err := h.QuestionService.GetQuestions(&user)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}
// Получить детали вопроса
func (h *Handler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	questionIDString := r.PathValue("questionID")
	if questionIDString == "" {
		http.Error(w, "missing questionID", http.StatusBadRequest)
	}
	questionID, _ := strconv.Atoi(questionIDString)

	versionString := r.PathValue("version")
	if versionString == "" {
		http.Error(w, "missing version", http.StatusBadRequest)
		return
	}
	version, _ := strconv.Atoi(versionString)

	question, err := h.QuestionService.GetQuestionByID(&user, questionID, version)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}
// Изменить вопрос
func (h *Handler) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	questionIDString := r.PathValue("questionID")
	if questionIDString == "" {
		http.Error(w, "missing questionID", http.StatusBadRequest)
	}
	questionID, _ := strconv.Atoi(questionIDString)

	var data map[string]any

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	title, hasTitle := data["title"].(string)
	content, hasContent := data["content"].(string)
	options, hasOptions := data["options"].([]string)
	correctOption, hasCorrectOptions := data["correct_option"].(int)


	err := h.QuestionService.UpdateQuestion(
		&user, 
		questionID, 
		title, hasTitle,
		content, hasContent,
		options, hasOptions,
		correctOption, hasCorrectOptions,
	)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}
// Создать вопрос
func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	var data map[string]any

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	title, hasTitle := data["title"].(string)
	content, hasContent := data["content"].(string)
	optionsInterface, hasOptions := data["options"].([]interface{})
	correctOptionFloat, hasCorrectOptions := data["correct_option"].(float64)

	if !hasTitle || !hasContent || !hasOptions || !hasCorrectOptions {
		http.Error(w, "not enough fields", 400)
		return
	}

    var options []string
    for _, opt := range optionsInterface {
        if str, ok := opt.(string); ok {
            options = append(options, str)
        } else {
            http.Error(w, "options must be strings", 400)
            return
        }
    }

    correctOption := int(correctOptionFloat)

    if correctOption < 0 || correctOption >= len(options) {
        http.Error(w, "correct_option out of range", 400)
        return
    }

	questionData, err := h.QuestionService.CreateQuestion(
		&user,
		title,
		content,
		options,
		correctOption,
	)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questionData)
}
// Удалить вопрос
func (h *Handler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	questionIDString := r.PathValue("questionID")
	if questionIDString == "" {
		http.Error(w, "missing questionID", http.StatusBadRequest)
	}
	questionID, _ := strconv.Atoi(questionIDString)

	err := h.QuestionService.DeleteQuestion(&user, questionID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}