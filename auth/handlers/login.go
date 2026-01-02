package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	"fmt"
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
	loginType := r.URL.Query().Get("type")
	if loginType != "code" {
		http.Error(w, "unsupported login type", http.StatusBadRequest)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EntryToken == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// 1. сохраняем состояние авторизации
	authState := domain.AuthState{
		EntryToken: req.EntryToken,
		Status:     domain.StatusPending,
		ExpiresAt:  time.Now().Add(5 * time.Minute),
	}
	storage.SaveAuthState(authState)

	// 2. генерируем код
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000

	codeState := domain.CodeState{
		Code: fmt.Sprintf("%06d", code),
		EntryToken: req.EntryToken,
		ExpiresAt:  time.Now().Add(1 * time.Minute),
	}
	storage.SaveCode(codeState)

	// 3. отдаём код клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CodeResponse{
		Code: codeState.Code,
	})
}
