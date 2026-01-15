package handlers

import (
	"encoding/json"
	"net/http"

	"auth/storage"
	"auth/domain"
)
func LoginCheck(w http.ResponseWriter, r *http.Request) {
	entryToken := r.URL.Query().Get("entry_token")
	if entryToken == "" {
		http.Error(w, "entry_token required", http.StatusBadRequest)
		return
	}

	state, ok := storage.GetAuthState(entryToken)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// если ещё не авторизован / отклонен
	if state.Status != domain.StatusApproved {
		json.NewEncoder(w).Encode(map[string]string{
			"status": string(state.Status),
		})
		return
	}

	// success
	json.NewEncoder(w).Encode(map[string]string{
		"status":        "approved",
		"access_token":  state.AccessToken,
		"refresh_token": state.RefreshToken,
        "user_id":       state.UserID,
	})
}





