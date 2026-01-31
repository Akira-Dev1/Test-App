package transport

import (
	"encoding/json"
	"net/http"
	"strconv"

	authTransport "main_module/internal/auth/transport"
)

// Получить список пользователей прошедших тест
func (h *Handler) GetPassedUsers(w http.ResponseWriter, r *http.Request) {
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

	studentsIDS, err := h.TestService.GetPassedUsers(&user, testID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentsIDS)
}
// Получить оценку пользователя
func (h *Handler) GetUserRatings(w http.ResponseWriter, r *http.Request) {
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


	scores, err := h.TestService.GetUserRatings(&user, testID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scores)
}
// Получить ответы пользователя
func (h *Handler) GetUserAnswers(w http.ResponseWriter, r *http.Request) {
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

	answers, err := h.TestService.GetUserAnswers(&user, testID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answers)
}
