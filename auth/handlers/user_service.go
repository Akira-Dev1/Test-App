package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"auth/user_service"
)

func UserService(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("type") == "get_user_list" {
		users, err := user_service.GetUserList()
		if err != nil {
		    http.Error(w, "no users in mongo :>", http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}

	if r.URL.Query().Get("type") == "get_user_info" {
		var req user_service.UserInfRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request :>", http.StatusBadRequest)
			return
		}
		log.Println(req.UserId)

		user, err := user_service.GetUserInfo(req.UserId)
		if err != nil {
		    http.Error(w, "no such user :>", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}