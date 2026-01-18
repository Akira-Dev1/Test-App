package handlers

import (
	"net/http"

	"auth/services"
	"auth/storage"
	"auth/domain"
	"auth/jwt"
)


func GithubCallback(w http.ResponseWriter, r *http.Request) {
    state := r.URL.Query().Get("state")
    code  := r.URL.Query().Get("code")
    errQ  := r.URL.Query().Get("error")

    if state == "" {
        http.Error(w, "state not found", http.StatusBadRequest)
        return
    }

    authState, ok := storage.GetAuthState(state)
    if !ok {
        http.Error(w, "invalid state", http.StatusBadRequest)
        return
    }

    if errQ != "" {
        authState.Status = domain.StatusDenied
        storage.SaveAuthState(authState, state)
        return
    }

    user, err := services.GithubAuth(code)
    if err != nil {
        http.Error(w, "github auth failed", http.StatusUnauthorized)
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

    storage.SaveAuthState(authState, state)

	

    w.Write([]byte("GitHub auth success"))
}

