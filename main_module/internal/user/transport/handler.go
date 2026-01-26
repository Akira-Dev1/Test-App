package transport

import (
	"encoding/json"
	"net/http"

	"main_module/internal/user/application"
)

type Handler struct {
	UserService *application.UserService
}

func NewUserHandler(userService *application.UserService) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

