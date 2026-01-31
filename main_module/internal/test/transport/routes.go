package transport

import (
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

func RegisterTestRoutes(mux *http.ServeMux, h *Handler) {
	//test_composition
	mux.Handle("DELETE /tests/{testID}/questions/{questionID}", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.RemoveQuestionFromTest))) 	// Удалить вопрос из теста
	mux.Handle("POST /tests/{testID}/questions/{questionID}", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.AddQuestionToTest))) 		// Добавить вопрос в тест
	mux.Handle("PUT /tests/{testID}/questions/reorder", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.ChangeOrderOfQuestions))) 	// Изменить порядок следования вопросов в тесте
	mux.Handle("GET /tests/{testID}/questions", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetTestQuestions))) 		// Получить список вопросов в тесте

	// test_results	
	mux.Handle("GET /tests/{testID}/passed-users", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetPassedUsers))) 			// Получить список пользователей прошедших тест
	mux.Handle("GET /tests/{testID}/scores", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetUserRatings))) 			// Получить оценку пользователя
	mux.Handle("GET /tests/{testID}/answers", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetUserAnswers))) 			// Получить ответы пользователя
}