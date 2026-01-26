package transport

import (
	"encoding/json"
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

// Получение информации о курсе (Список студентов)
func (h *Handler) GetCourseStudents(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	courseID := r.PathValue("courseID")
	if courseID == "" {
		http.Error(w, "missing courseID", http.StatusBadRequest)
	}

	students, err := h.CourseService.GetCourseStudents(&user, courseID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// Запись пользователя на курс
func (h *Handler) EnrollUser(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	courseID := r.PathValue("courseID")
	if courseID == "" {
		http.Error(w, "missing courseID", http.StatusBadRequest)
	}

	var data map[string]any

    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

	targetUserID, hasTargetUserID:= data["user_id"].(string)

	if !hasTargetUserID {
		http.Error(w, "required field is not enough", 400)
		return
	}

	if err := h.CourseService.EnrollUser(&user, courseID, targetUserID); err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}

// Отчисление пользователя с курса
func (h *Handler) UnenrollUser(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	courseID := r.PathValue("courseID")
	if courseID == "" {
		http.Error(w, "missing courseID", http.StatusBadRequest)
	}

	var data map[string]any

    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

	targetUserID, hasTargetUserID := data["user_id"].(string)

	if !hasTargetUserID {
		http.Error(w, "required field is not enough", 400)
		return
	}

	if err := h.CourseService.UnenrollUser(&user, courseID, targetUserID); err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}