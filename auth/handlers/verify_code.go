package handlers

import (
	"auth/codeauth"
	"auth/domain"
	"auth/jwt"
	"auth/storage"

	"encoding/json"
	"net/http"
)

func VerifyCode(w http.ResponseWriter, r *http.Request) {
	var req domain.VerifyCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	status, err := codeauth.VerifyCode(req.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	codeState, _ := storage.GetCode(req.Code)

	authState, ok := storage.GetAuthState(codeState.EntryToken)
    if !ok {
        http.Error(w, "invalid state", http.StatusBadRequest)
        return
    }

	// Парсим токен
	claims, err := jwt.ParseRefreshToken(req.RefreshToken)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	var user *domain.User

	user, err = storage.FindUserByEmail(claims.Email) 
	if err != nil {
		http.Error(w, "user via such email not found", http.StatusInternalServerError)
		return 
	}

	accessToken, err := jwt.GenerateAccessToken(
        user.ID.Hex(),
        user.Permissions,
    )
    if err != nil {
        http.Error(w, "access token error", http.StatusInternalServerError)
        return
    }

    refreshToken, err := jwt.GenerateRefreshToken(user.Email)
    if err != nil {
        http.Error(w, "refresh token error", http.StatusInternalServerError)
        return
    }

    if err := storage.AddRefreshToken(user.ID, refreshToken); err != nil {
        http.Error(w, "save refresh token error", http.StatusInternalServerError)
        return
    }

    authState.Status = domain.StatusApproved
    authState.UserID = user.ID.Hex()
    authState.AccessToken = accessToken
    authState.RefreshToken = refreshToken

    storage.SaveAuthState(authState, codeState.EntryToken)


	json.NewEncoder(w).Encode(domain.VerifyCodeResponse{
		Status: string(status),
	})
}
