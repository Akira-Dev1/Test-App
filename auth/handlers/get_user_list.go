package handlers

import (
	"net/http"
	"encoding/json"

	"auth/user_service"
)

func GetUserList(w http.ResponseWriter, r *http.Request) () {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	users, err := user_service.GetUserList()
	if err != nil {
    http.Error(w, "no users in mongo :<", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}