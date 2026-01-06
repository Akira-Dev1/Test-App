package handlers

import (
	"net/http"
	"log"
)

func GithubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code not provided", http.StatusBadRequest)
		return
	}
	log.Println("GitHub code:", code)

	// дальше:
	// 1. обмен code → access_token
	// 2. запрос user info
	// 3. работа с Mongo
}
