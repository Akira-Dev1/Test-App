package handlers

import (
	"net/http"
	"encoding/json"

	"auth/user_service"
)

func GetUserBlockStatus (w http.ResponseWriter, r *http.Request) () {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req user_service.UserGetBlockStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request :<", http.StatusBadRequest)
		return
	}

	status, err := user_service.GetUserBlockStatus(req.UserId)
	if err != nil {
	    http.Error(w, "no such user :<", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}