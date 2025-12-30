package main

import (
	"log"
	"net/http"

	"auth/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/check", handler.Check)
	mux.HandleFunc("/logout", handler.Logout)

	log.Println("Auth service started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

