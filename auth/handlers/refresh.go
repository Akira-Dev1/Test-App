package handlers

import(
	"net/http"

	"auth/jwt"
	"auth/permissions"
	"encoding/json"
	"auth/storage"
)

type RefreshResponse struct{
	RefreshToken string `json:"refresh_token"`
}


func contains(tokens []string, target string) bool {
	for _, t := range tokens {
		if t == target {
			return true
		}
	}
	return false
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// 1. Извлекаем токен
	var req RefreshResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request :<", http.StatusBadRequest)
		return
	}

	oldRefresh := req.RefreshToken

	// 2. Проверяем JWT refresh
	claims, err := jwt.ParseRefreshToken(oldRefresh)
	if err != nil {
		http.Error(w, "invalid refresh token", http.StatusUnauthorized)
		return
	}

	// 3. Ищем пользователя
	user, err := storage.FindUserByEmail(claims.Email)
	if err != nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}

	// 4. Проверяем, что refresh реально наш
	if !contains(user.RefreshTokens, oldRefresh) {
		http.Error(w, "refresh revoked", http.StatusUnauthorized)
		return
	}

	// 5. УДАЛЯЕМ старый refresh (rotation)
	if err := storage.RemoveRefreshToken(user.ID, oldRefresh); err != nil {
		http.Error(w, "failed to rotate refresh", http.StatusInternalServerError)
		return
	}

	// 6. Генерируем новые токены
	perms := permissions.ResolvePermissions(user.Roles)

	accessToken, err := jwt.GenerateAccessToken(
		user.ID.Hex(),
		perms,
	)
	if err != nil {
		http.Error(w, "failed to generate access", http.StatusInternalServerError)
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.Email)
	if err != nil {
		http.Error(w, "failed to generate refresh", http.StatusInternalServerError)
		return
	}

	// 7. Сохраняем новый refresh
	if err := storage.AddRefreshToken(user.ID, refreshToken); err != nil {
		http.Error(w, "failed to save refresh", http.StatusInternalServerError)
		return
	}

	// 8. Ответ
	resp := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
