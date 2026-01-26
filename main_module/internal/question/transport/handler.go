package transport

import (
	"encoding/json"
	"net/http"

	"main_module/internal/question/application"
)

type Handler struct {
	QuestionService *application.QuestionService
}

func NewQuestionHandler(questionService *application.QuestionService) *Handler {
	return &Handler{
		QuestionService: questionService,
	}
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

