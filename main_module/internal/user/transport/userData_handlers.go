package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	authTransport "main_module/internal/auth/transport"
)

// Получить информацию о пользователе (курсы, оценки, тесты)
func (h *Handler) GetUserData(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	targetUserID := r.PathValue("targetUserID")
	if targetUserID == "" {
		http.Error(w, "missing targetUserID", http.StatusBadRequest)
	}

    query := r.URL.Query()
    
    // Парсим с дефолтным значением false если не передано
    hasCourses, _ := strconv.ParseBool(query.Get("courses"))
    hasTests, _ := strconv.ParseBool(query.Get("tests"))
    hasGrades, _ := strconv.ParseBool(query.Get("grades"))

	userData, err := h.UserService.GetUserData(&user, targetUserID, hasCourses, hasTests, hasGrades)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}