package transport

import (
	"encoding/json"
	"net/http"

	authTransport "main_module/internal/auth/transport"
	"main_module/internal/course/application"
)

type Handler struct {
	CourseService *application.CourseService
}

func NewCourseHandler(courseService *application.CourseService) *Handler {
	return &Handler{
		CourseService: courseService,
	}
}

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

	json.NewEncoder(w).Encode(courses)
}
