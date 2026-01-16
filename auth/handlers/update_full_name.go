package handlers

import (
	"net/http"
	"encoding/json"

	"auth/user_service"
)

func UpdateFullName (w http.ResponseWriter, r *http.Request) () {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req user_service.UpdateUserNameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request :<", http.StatusBadRequest)
		return
	}

	if err := user_service.UpdateUserName(req.UserId, req.NewName); err != nil {
		http.Error(w, "Couldn't update user :<", http.StatusNotFound)
		return
	}
	w.Write([]byte("Username successfully updated!"))
}