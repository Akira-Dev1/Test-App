package codeauth

import (
	"fmt"
	"math/rand"
	"time"

	"auth/domain"
	"auth/storage"
)

func GenerateCode(entryToken string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(900000)+100000)

	codeState := domain.CodeState{
		EntryToken: entryToken,
		ExpiresAt:  time.Now().Add(1 * time.Minute),
	}
	storage.SaveCode(codeState, code)

	return code, nil
}
