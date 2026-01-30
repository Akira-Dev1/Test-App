package transport

import (
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

func RegisterQuestionRoutes(mux *http.ServeMux, h *Handler) {
	mux.Handle("GET /questions", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetQuestions))) 			// Получить список вопросов
	mux.Handle("GET /questions/{questionID}/{version}", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetQuestionByID))) 			// Получить детали вопроса
	mux.Handle("POST /questions/{questionID}", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.UpdateQuestion))) 			// Изменить вопрос
	mux.Handle("POST /questions", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.CreateQuestion))) 			// Создать вопрос
	mux.Handle("DELETE /questions/{questionID}", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.DeleteQuestion))) 			// Удалить вопрос
}