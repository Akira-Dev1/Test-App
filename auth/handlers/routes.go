package handlers

import "net/http"

func RegisterRoutes() {

	
	http.HandleFunc("/login", Login)
	http.HandleFunc("/login/check", LoginCheck)
}
