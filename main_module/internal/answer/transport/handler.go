package transport

import (
	"encoding/json"
	"net/http"

	"main_module/internal/answer/application"
)

type Handler struct {
	AnswerService *application.AnswerService
}

func NewAnswerHandler(answerService *application.AnswerService) *Handler {
	return &Handler{
		AnswerService: answerService,
	}
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

