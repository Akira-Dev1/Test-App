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
	) // Получение списка курсов
	mux.Handle(
		"GET /courses/{id}",
		authTransport.AuthMiddleware(
			http.HandlerFunc(h.GetCourseByID),
		),
	) // Получение курса по айди
	mux.Handle(
		"PATCH /courses/{id}",
		authTransport.AuthMiddleware(
			http.HandlerFunc(h.UpdateCourse),
		),
	) // Изменение курса
	mux.Handle(
		"POST /courses",
		authTransport.AuthMiddleware(
			http.HandlerFunc(h.CreateCourse),
		),
	) // Создание курса
	mux.Handle(
		"DELETE /courses/{id}",
		authTransport.AuthMiddleware(
			http.HandlerFunc(h.DeleteCourse),
		),
	) // Удаление курса
}