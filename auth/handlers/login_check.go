package handlers

import (
	"encoding/json"
	"net/http"

	"auth/storage"
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
	json.NewEncoder(w).Encode(map[string]string{
		"status": string(state.Status),
	})
}
