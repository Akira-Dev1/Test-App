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

// Функция парсинга refresh токена
func ParseRefreshToken(tokenString string) (*RefreshClaims, error) {
    // 1. Парсим токен с проверкой подписи
    token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
        // 2. Проверяем алгоритм подписи
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
        }
        // 3. Возвращаем секретный ключ для проверки
        return config.JWTRefreshSecret, nil
    })

    // 4. Если ошибка парсинга
    if err != nil {
        return nil, fmt.Errorf("ошибка парсинга токена: %w", err)
    }

    // 5. Извлекаем claims и проверяем валидность
    if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
        return claims, nil
    }

    // 6. Если токен невалиден
    return nil, fmt.Errorf("невалидный токен")
}