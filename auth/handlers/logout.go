package handlers

import (
	"net/http"
	"encoding/json"

	"auth/jwt"
	"auth/storage"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("all") == "true"{
		var req RefreshResponse

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request :<", http.StatusBadRequest)
			return
		}

		refresh := req.RefreshToken

		claims, err := jwt.ParseRefreshToken(refresh)
		if err != nil {
			http.Error(w, "invalid refresh token", http.StatusUnauthorized)
			return
		}

		user, err := storage.FindUserByEmail(claims.Email)
		if err != nil {
			http.Error(w, "user not found", http.StatusUnauthorized)
			return
		}

		if err := storage.RemoveAllRefreshTokens(user.ID); err != nil {
			http.Error(w, "failed to logout", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok (logged out successfully EVERYWHERE)",
		})
	} else {
		var req RefreshResponse

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request :<", http.StatusBadRequest)
			return
		}

		refresh := req.RefreshToken

		claims, err := jwt.ParseRefreshToken(refresh)
		if err != nil {
			http.Error(w, "invalid refresh token", http.StatusUnauthorized)
			return
		}

		user, err := storage.FindUserByEmail(claims.Email)
		if err != nil {
			http.Error(w, "user not found", http.StatusUnauthorized)
			return
		}

		if err := storage.RemoveRefreshToken(user.ID, refresh); err != nil {
			http.Error(w, "failed to logout", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok (logged out successfully)",
		})
	}
}
