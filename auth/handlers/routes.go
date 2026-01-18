package handlers

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/login/check", LoginCheck)
	http.HandleFunc("/login/verify", VerifyCode)
	http.HandleFunc("/login/github/callback", GithubCallback)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/yandex/callback", YandexCallback)


	http.HandleFunc("/userservice", UserService)

	http.HandleFunc("/userservice/get_user_list", GetUserList)
	http.HandleFunc("/userservice/get_user_block_status", GetUserBlockStatus)
	http.HandleFunc("/userservice/get_user_info", GetUserInfo)
	http.HandleFunc("/userservice/get_user_roles", GetUserRoles)

	http.HandleFunc("/userservice/set_block_status", SetBlockStatus)
	
	http.HandleFunc("/userservice/update_full_name", UpdateFullName)
	http.HandleFunc("/userservice/update_user_roles", UpdateUserRoles)
}
