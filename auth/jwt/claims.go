package jwt

import "github.com/golang-jwt/jwt/v5"

type AccessClaims struct {
	UserID      string   `json:"user_id"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
