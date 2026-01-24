package transport

import (
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

func RegisterCourseRoutes(mux *http.ServeMux, h *Handler) {
	mux.Handle(
		"GET /courses",
		authTransport.AuthMiddleware(
			http.HandlerFunc(h.GetCourses),
		),
	)
}