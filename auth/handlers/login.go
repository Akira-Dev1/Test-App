package handlers

import (
	"encoding/json"
	"net/http"
	"auth/codeauth"
	"time"

	"auth/domain"
	"auth/storage"
)

type LoginRequest struct {
	EntryToken string `json:"entry_token"`
}

type CodeResponse struct {
	Code string `json:"code"`
}


func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("type") != "code" {
		http.Error(w, "unsupported login type", http.StatusBadRequest)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EntryToken == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	// генерируем код
	code, err := codeauth.GenerateCode(req.EntryToken) // здесь респонс от компонента Code Autentification - он возвр. код

	// создаём auth state
	authState := domain.AuthState{
		Status:     domain.StatusPending,
		ExpiresAt:  time.Now().Add(5 * time.Minute),
	}
	storage.SaveAuthState(authState, req.EntryToken)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(CodeResponse{Code: code})
}

