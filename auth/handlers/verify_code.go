package handlers

import (
	"encoding/json"
	"net/http"
	"auth/domain"
	"auth/codeauth"
)

func VerifyCode(w http.ResponseWriter, r *http.Request) {
	var req domain.VerifyCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	status, err := codeauth.VerifyCode(req.EntryToken, req.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(domain.VerifyCodeResponse{
		Status: string(status),
	})
}


