package infrastructure

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct{}


type Claims struct {
    UserID      string      `json:"user_id"`
    Permissions []string `json:"permissions"`
    Blocked     bool     `json:"blocked"`
	jwt.RegisteredClaims
}

func ParseToken(tokenString string) (*Claims, error) {
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secret) == 0 {
		return nil, errors.New("jwt secret not set in env")
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims := token.Claims.(*Claims)

	if claims.Blocked {
		return nil, errors.New("user blocked")
	}

	return claims, nil

}