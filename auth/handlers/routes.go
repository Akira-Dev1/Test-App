package handlers

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/login/check", LoginCheck)
	http.HandleFunc("/login/verify", VerifyCode)
	http.HandleFunc("/login/github/callback", GithubCallback)
	http.HandleFunc("/refresh", Refresh)
	// logout
	http.HandleFunc("/yandex/callback", YandexCallback)
	http.HandleFunc("/userservice", UserService)
}
