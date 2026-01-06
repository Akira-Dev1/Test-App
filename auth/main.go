package main

import (
	"log"
	"net/http"
	
	"auth/config"
	"auth/storage"
	"auth/handlers"
)

func main() {
	config.Connect("mongodb://mongo:27017")
	storage.InitUserCollection("authdb", "users")

	handlers.RegisterRoutes()

	log.Println("Auth service started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
