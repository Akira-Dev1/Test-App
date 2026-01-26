package app

import (
    "database/sql"
	"net/http"
	"log"

	courseApp 			"main_module/internal/course/application"
	courseInfra 		"main_module/internal/course/infrastructure"
	courseTransport 	"main_module/internal/course/transport"

	testApp 			"main_module/internal/test/application"
	testInfra 			"main_module/internal/test/infrastructure"
	testTransport 		"main_module/internal/test/transport"

	questionApp 		"main_module/internal/question/application"
	questionInfra 		"main_module/internal/question/infrastructure"
	questionTransport 	"main_module/internal/question/transport"

	attemptApp 			"main_module/internal/attempt/application"
	attemptInfra 		"main_module/internal/attempt/infrastructure"
	attemptTransport 	"main_module/internal/attempt/transport"

	userApp				"main_module/internal/user/application"
	userInfra			"main_module/internal/user/infrastructure"
	userTransport		"main_module/internal/user/transport"

	answerApp			"main_module/internal/answer/application"
	answerInfra			"main_module/internal/answer/infrastructure"
	answerTransport		"main_module/internal/answer/transport"
)

type App struct {
    mux *http.ServeMux
}

func NewApp(db *sql.DB) *App {
	courseRepo := courseInfra.NewCourseRepository(db)
	courseService := courseApp.NewCourseService(courseRepo)
	courseHandler := courseTransport.NewCourseHandler(courseService)

	testRepo := testInfra.NewTestRepository(db)
	testService := testApp.NewTestService(testRepo)
	testHandler := testTransport.NewTestHandler(testService)

	questionRepo := questionInfra.NewQuestionRepository(db)
	questionService := questionApp.NewQuestionService(questionRepo)
	questionHandler := questionTransport.NewQuestionHandler(questionService)

	attemptRepo := attemptInfra.NewAttemptRepository(db)
	attemptService := attemptApp.NewAttemptService(attemptRepo)
	attemptHandler := attemptTransport.NewAttemptHandler(attemptService)

	userRepo := userInfra.NewUserRepository(db)
	userService := userApp.NewUserService(userRepo)
	userHandler := userTransport.NewUserHandler(userService)

	answerRepo := answerInfra.NewAnswerRepository(db)
	answerService := answerApp.NewAnswerService(answerRepo)
	answerHandler := answerTransport.NewAnswerHandler(answerService)



	mux := http.NewServeMux()

	courseTransport.RegisterCourseRoutes(mux, courseHandler)
	testTransport.RegisterTestRoutes(mux, testHandler)
	questionTransport.RegisterQuestionRoutes(mux, questionHandler)
	attemptTransport.RegisterAttemptRoutes(mux, attemptHandler)
	userTransport.RegisterUserRoutes(mux, userHandler)
	answerTransport.RegisterAnswerRoutes(mux, answerHandler)

    return &App{mux: mux}
}

func (a *App) Run() {
    log.Println("Server started on :18080")
    http.ListenAndServe(":18080", a.mux)
}
