package jwt

import (
	"time"

	"auth/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID string, permissions []string) (string, error) {
	claims := AccessClaims{
		UserID:      userID,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTAccessSecret)
}
