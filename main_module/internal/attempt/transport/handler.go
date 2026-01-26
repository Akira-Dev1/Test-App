package transport

import (
	"encoding/json"
	"net/http"

	"main_module/internal/attempt/application"
)

type Handler struct {
	AttemptService *application.AttemptService
}

func NewAttemptHandler(attemptService *application.AttemptService) *Handler {
	return &Handler{
		AttemptService: attemptService,
	}
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

