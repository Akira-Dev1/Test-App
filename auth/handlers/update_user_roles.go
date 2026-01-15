package handlers

import (
	"encoding/json"
	"net/http"

	"auth/user_service"
)

func UpdateUserRoles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req user_service.UpdateUserRolesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request :<", http.StatusBadRequest)
		return
	}

	if err := user_service.UpdateUserRoles(req.UserId, req.Roles); err != nil {
		http.Error(w, "Couldn't update user :<", http.StatusNotFound)
		return
	}
	w.Write([]byte("Roles successfully updated!"))
}
