package transport

import (
	"encoding/json"
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

// Получить список пользователей
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	targetUserID := r.PathValue("targetUserID")
	if targetUserID == "" {
		http.Error(w, "missing targetUserID", http.StatusBadRequest)
	}

	users, err := h.UserService.GetUsers(&user)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Получить информацию о пользователе (Список ролей)
func (h *Handler) GetUserRoles(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	targetUserID := r.PathValue("targetUserID")
	if targetUserID == "" {
		http.Error(w, "missing targetUserID", http.StatusBadRequest)
	}

	userData, err := h.UserService.GetUserRoles(&user, targetUserID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}

// Изменить информацию о пользователе (Список ролей)
func (h *Handler) UpdateUserRoles(w http.ResponseWriter, r *http.Request) {
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

	roles, hasRoles := data["roles"].([]string)

	if !hasRoles {
		http.Error(w, "the roles field is missing", 400)
		return
	}

	err := h.UserService.UpdateUserRoles(&user, targetUserID, roles)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}

// Получить информацию о пользователе (Заблокирован ли пользователь)
func (h *Handler) GetUserStatus(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	targetUserID := r.PathValue("targetUserID")
	if targetUserID == "" {
		http.Error(w, "missing targetUserID", http.StatusBadRequest)
	}

	userData, err := h.UserService.GetUserStatus(&user, targetUserID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}

// Изменить информацию о пользователе (Заблокировать/Разблокировать пользователя)
func (h *Handler) UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
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

	status, hasStatus := data["is_blocked"].(bool)

	if !hasStatus {
		http.Error(w, "the is_blocked field is missing", 400)
		return
	}
	err := h.UserService.UpdateUserStatus(&user, targetUserID, status)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}