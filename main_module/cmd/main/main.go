package main

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"net/http"
	"log"

	courseApp 		"main_module/internal/course/application"
	courseInfra 	"main_module/internal/course/infrastructure"
	courseTransport "main_module/internal/course/transport"
)


func main() {
	db, err := sql.Open(
		"postgres", 
		"host=postgres port=5432 user=core password=core dbname=core_db sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	courseRepo := courseInfra.NewCourseRepository(db)
	courseService := courseApp.NewCourseService(courseRepo)
	courseHandler := courseTransport.NewCourseHandler(courseService)

	mux := http.NewServeMux()
	courseTransport.RegisterCourseRoutes(mux, courseHandler)

	log.Println("started :18080")
	http.ListenAndServe(":18080", mux)
}