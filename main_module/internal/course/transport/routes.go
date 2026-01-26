package transport

import (
	"net/http"

	authTransport "main_module/internal/auth/transport"
)

func RegisterCourseRoutes(mux *http.ServeMux, h *Handler) {
	// Course
	mux.Handle("GET /courses", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetCourses))) 			// Получение списка курсов
	mux.Handle("GET /courses/{courseID}", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetCourseByID))) 		// Получение информации о курсе (Название, Описание, ID преподавателя)
	mux.Handle("PATCH /courses/{courseID}", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.UpdateCourse))) 		// Изменение курса
	mux.Handle("POST /courses", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.CreateCourse))) 		// Создание курса
	mux.Handle("DELETE /courses/{courseID}", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.DeleteCourse))) 		// Удаление курса

	// Test
	mux.Handle("GET /courses/{courseID}/tests", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetCourseTests)))		// Получение информации о курсе (Список тестов)
	mux.Handle("GET /courses/{courseID}/tests/{testID}/status", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetTestStatus)))		// Получение информации о тесте (Активный тест или нет)
	mux.Handle("PATCH /courses/{courseID}/tests/{testID}/activation", 
		authTransport.AuthMiddleware(http.HandlerFunc(h.UpdateTestStatus)))		// Изменение теста (Активировать/Деактивировать тест)
	mux.Handle("POST /courses/{courseID}/tests",
		authTransport.AuthMiddleware(http.HandlerFunc(h.CreateTest)))			// Создание теста
	mux.Handle("DELETE /courses/{courseID}/tests/{testID}",
		authTransport.AuthMiddleware(http.HandlerFunc(h.DeleteTest)))			// Удаление теста

	// Enrollment
	mux.Handle("GET /courses/{courseID}/students",
		authTransport.AuthMiddleware(http.HandlerFunc(h.GetCourseStudents)))	// Получение информации о курсе (Список студентов)
	mux.Handle("POST /courses/{courseID}/join",
		authTransport.AuthMiddleware(http.HandlerFunc(h.EnrollUser)))			// Запись пользователя на курс
	mux.Handle("DELETE /courses/{courseID}/leave",	
		authTransport.AuthMiddleware(http.HandlerFunc(h.UnenrollUser)))			// Отчисление пользователя с курса
}