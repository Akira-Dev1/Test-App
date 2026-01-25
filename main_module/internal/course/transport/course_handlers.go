package transport

import (
	"encoding/json"
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

// Получить список курсов
func (h *Handler) GetCourses(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	courses, err := h.CourseService.GetCourses(&user)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Получить курс по айди
func (h *Handler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
	}

	course, err := h.CourseService.GetCourseByID(&user, id)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

// Изменить курс
func (h *Handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
	}

	var data map[string]any

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	var title any = data["title"]
	var description any= data["description"]

	if title == nil && description == nil {
		http.Error(w, "no fields to update", 400)
		return
	}

	if err := h.CourseService.UpdateCourse(&user, id, title, description); err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}

// Создать курса
func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	var data map[string]any

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	title, hasTitle := data["title"].(string)
	description, hasDescription := data["description"].(string)

	if !hasTitle || !hasDescription {
		http.Error(w, "not enought fields to create", 400)
		return
	}

	id, err := h.CourseService.CreateCourse(&user, title, description)
	if err != nil{
		http.Error(w, err.Error(), 403)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

// Удалить курса
func (h *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	user, ok := authTransport.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", 401)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
	}

	if err := h.CourseService.DeleteCourse(&user, id); err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
}