package codeauth

import (
	"errors"
	"time"
	"auth/domain"
	"auth/storage"
)

func VerifyCode(code string) (domain.AuthStatus, error) {
	codeState, ok := storage.GetCode(code)
	if !ok {
		return domain.StatusDenied, ErrCodeNotFound
	}

	if time.Now().After(codeState.ExpiresAt) {
		deny(codeState.EntryToken)
		return domain.StatusDenied, ErrCodeExpired
	}

	authState, ok := storage.GetAuthState(codeState.EntryToken)
	if !ok {
		return domain.StatusDenied, errors.New("auth state not found")
	}

	authState.Status = domain.StatusApproved
	storage.SaveAuthState(authState, codeState.EntryToken)

	return authState.Status, nil
}

func deny(entryToken string) {
	if authState, ok := storage.GetAuthState(entryToken); ok {
		authState.Status = domain.StatusDenied
		storage.SaveAuthState(authState, entryToken)
	}
}
