package handlers

import (
	"net/http"
)

func RegisterRoutes(){
	// Авторизация
	http.HandleFunc("/get_user_status", GetUserStatus)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/login/check", UpdateUserStatusHandler)
	http.HandleFunc("/login/verify", VerifyCode)

	// Уведомления
	http.HandleFunc("/notifications", GetNotifications)

	// Курсы
    http.HandleFunc("/courses/list", GetCoursesHandler)
    http.HandleFunc("/courses/get", GetCourseHandler)
    http.HandleFunc("/courses/create", CreateCourseHandler)
	http.HandleFunc("/courses/update", UpdateCourseHandler)
    http.HandleFunc("/courses/delete", DeleteCourseHandler)
    http.HandleFunc("/courses/join", JoinCourseHandler) 
	http.HandleFunc("/courses/leave", LeaveCourseHandler) 

	// Тесты
    http.HandleFunc("/tests", GetTestsHandler)
    http.HandleFunc("/tests/create", CreateTestHandler)
	http.HandleFunc("/tests/delete", DeleteTestHandler)
    http.HandleFunc("/tests/status", GetTestStatusHandler)
    http.HandleFunc("/tests/activation", ToggleTestActivationHandler)

	// Вопросы в тестах
    http.HandleFunc("/tests/questions/reorder", ReorderQuestionsHandler)
    http.HandleFunc("/tests/questions/add", AddQuestionToTestHandler)
    http.HandleFunc("/tests/questions/remove", RemoveQuestionFromTestHandler)

	// Вопросы
    http.HandleFunc("/questions/create", CreateQuestionHandler)
    http.HandleFunc("/questions/my", GetMyQuestionsHandler)
    http.HandleFunc("/questions/update", UpdateQuestionHandler)
	http.HandleFunc("/questions/delete", DeleteQuestionHandler)
    http.HandleFunc("/questions/detail", GetQuestionDetailHandler)

	// Попытки
	http.HandleFunc("/tests/passed", GetPassedUsersHandler)
    http.HandleFunc("/tests/scores", GetScoresHandler)
    http.HandleFunc("/tests/answers", GetAnswersHandler)
    http.HandleFunc("/tests/start", StartAttemptHandler)
    http.HandleFunc("/attempts/answers/send", SubmitAnswerHandler)
    http.HandleFunc("/attempts/answers/update", UpdateAnswerHandler)
	http.HandleFunc("/attempts/answers/delete", DeleteAnswerHandler)
    http.HandleFunc("/attempts/complete", CompleteAttemptHandler)
    http.HandleFunc("/tests/user/attempt", GetUserAttemptHandler)
    http.HandleFunc("/tests/user/answers", GetUserAttemptAnswersHandler)

	// Пользователи
    http.HandleFunc("/users/profile", GetUserDataProfileHandler)
    http.HandleFunc("/users/list", GetUsersListHandler)
	http.HandleFunc("/users/info", GetUserInfoHandler)
    http.HandleFunc("/users/update_name", UpdateFullNameHandler)
    http.HandleFunc("/users/roles", GetUserRolesHandler)
	http.HandleFunc("/users/roles/update", UpdateUserRolesHandler)
    http.HandleFunc("/users/block/status", GetUserBlockStatusHandler)
    http.HandleFunc("/users/block/set", SetUserBlockHandler)
}
