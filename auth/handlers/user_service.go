package handlers

import (
	"encoding/json"
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
}