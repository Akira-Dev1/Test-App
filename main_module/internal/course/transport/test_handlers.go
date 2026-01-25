package transport

import (
	"encoding/json"
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

// Получение информации о курсе (Список тестов)
func (h *Handler) GetCourseTests(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	courseID := r.PathValue("courseID")
	if courseID == "" {
		http.Error(w, "missing courseID", http.StatusBadRequest)
	}

	tests, err := h.CourseService.GetCourseTests(&user, courseID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tests)
}

// Получение информации о тесте (Активный тест или нет)
func (h *Handler) GetTestStatus(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	courseID := r.PathValue("courseID")
	if courseID == "" {
		http.Error(w, "missing courseID", http.StatusBadRequest)
	}
	testID := r.PathValue("testID")
	if testID == "" {
		http.Error(w, "missing testID", http.StatusBadRequest)
	}
	status, err := h.CourseService.GetTestStatus(&user, courseID, testID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// Изменение теста (Активировать/Деактивировать тест)
func (h *Handler) UpdateTestStatus(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	courseID := r.PathValue("courseID")
	if courseID == "" {
		http.Error(w, "missing courseID", http.StatusBadRequest)
	}
	testID := r.PathValue("testID")
	if testID == "" {
		http.Error(w, "missing testID", http.StatusBadRequest)
	}

	var data map[string]any

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	var status any = data["status"]

	err := h.CourseService.UpdateTestStatus(&user, courseID, testID, status.(bool))
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}

// Создание теста
func (h *Handler) CreateTest(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "invalid json", 400)
		return
	}

	var title any = data["title"]

	testID, err := h.CourseService.CreateTest(&user, courseID, title.(string))
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testID)
}

// Удаление теста
func (h *Handler) DeleteTest(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	courseID := r.PathValue("courseID")
	if courseID == "" {
		http.Error(w, "missing courseID", http.StatusBadRequest)
	}
	testID := r.PathValue("testID")
	if testID == "" {
		http.Error(w, "missing testID", http.StatusBadRequest)
	}
	err := h.CourseService.DeleteTest(&user, courseID, testID)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}
