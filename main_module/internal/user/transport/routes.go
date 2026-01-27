package transport

import (
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

func RegisterUserRoutes(mux *http.ServeMux, h *Handler) {
	// Profile
	mux.Handle("GET /users/{targetUserID}/name", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetUserName))) 			// Получить информацию о пользователе (ФИО)
	mux.Handle("PATCH /users/{targetUserID}/name", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.UpdateUserName))) 		// Изменить информацию о пользователе (ФИО)

	// UserData
	mux.Handle("GET /users/{targetUserID}/data", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetUserData))) 			// Получить информацию о пользователе (курсы, оценки, тесты)

	//AdminPanel
	mux.Handle("GET /users", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetUsers))) 			// Получить список пользователей
	mux.Handle("GET /users/{targetUserID}/roles", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetUserRoles))) 		// Получить информацию о пользователе (Список ролей)
	mux.Handle("PATCH /users/{targetUserID}/roles", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.UpdateUserRoles))) 		// Изменить информацию о пользователе (Список ролей)
	mux.Handle("GET /users/{targetUserID}/block-status", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetUserStatus))) 		// Получить информацию о пользователе (Заблокирован ли пользователь)
	mux.Handle("PATCH /users/{targetUserID}/block", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.UpdateUserStatus))) 	// Изменить информацию о пользователе (Заблокировать/Разблокировать пользователя)
}