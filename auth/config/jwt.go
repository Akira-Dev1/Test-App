package config

import (
	"os"
)

func MustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("env var not set: " + key)
	}
	return val
}

var (
	JWTAccessSecret  = []byte(MustGetEnv("JWT_ACCESS_SECRET"))
	JWTRefreshSecret = []byte(MustGetEnv("JWT_REFRESH_SECRET"))
)
