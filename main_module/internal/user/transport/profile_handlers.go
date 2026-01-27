package transport

import (
	"encoding/json"
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

// Получить информацию о пользователе (ФИО)
func (h *Handler) GetUserName(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	targetUserID := r.PathValue("targetUserID")
	if targetUserID == "" {
		http.Error(w, "missing targetUserID", http.StatusBadRequest)
	}

	userData, err := h.UserService.GetUserName(&user, targetUserID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}

// Изменить информацию о пользователе (ФИО)
func (h *Handler) UpdateUserName(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	targetUserID := r.PathValue("targetUserID")
	if targetUserID == "" {
		http.Error(w, "missing targetUserID", http.StatusBadRequest)
	}

	var data map[string]any

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	name, hasName := data["full_name"].(string)

	if !hasName {
		http.Error(w, "the full_name field is missing", 400)
		return
	}

	err := h.UserService.UpdateUserName(&user, targetUserID, name)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}