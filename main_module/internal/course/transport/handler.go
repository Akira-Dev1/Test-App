package transport

import (
	"encoding/json"
	"net/http"

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

func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

