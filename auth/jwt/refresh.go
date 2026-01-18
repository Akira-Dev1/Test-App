package jwt

import (
	"time"
	"fmt"

	"auth/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(email string) (string, error) {
	claims := RefreshClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTRefreshSecret)
}

func ParseRefreshToken(tokenString string) (*RefreshClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
        }
        return config.JWTRefreshSecret, nil
    })

    if err != nil {
        return nil, fmt.Errorf("ошибка парсинга токена: %w", err)
    }

    if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("невалидный токен")
}