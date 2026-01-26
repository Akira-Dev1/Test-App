package transport

import (
	"encoding/json"
	"net/http"

	"main_module/internal/test/application"
)

type Handler struct {
	TestService *application.TestService
}

func NewTestHandler(testService *application.TestService) *Handler {
	return &Handler{
		TestService: testService,
	}
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

