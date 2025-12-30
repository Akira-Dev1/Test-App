package main

import (
	"auth/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/check", handlers.Check)
	mux.HandleFunc("/logout", handlers.Logout)
    mux.HandleFunc("/login/email/request", handlers.RequestLoginCode)

	log.Println("Auth service started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

