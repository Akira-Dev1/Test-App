package handlers

import (
	"net/http"
)

func RegisterRoutes(){
	http.HandleFunc("/get_user_status", GetUserStatus)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/login/check", UpdateUserStatusHandler)
}
