package main

import (
	"log"
	"net/http"

	"auth/handlers"
)

func main() {
	handlers.RegisterRoutes()

	log.Println("Auth service started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
