package main

import (
	"auth/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.InitRedis("redis:6379") 

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/check", handlers.Check)
	mux.HandleFunc("/logout", handlers.Logout)
    mux.HandleFunc("/login/email/request", handlers.RequestLoginCode)
	mux.HandleFunc("/login/email/verify", handlers.VerifyLoginCode)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

