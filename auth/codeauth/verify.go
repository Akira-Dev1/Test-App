package codeauth

import (
	"errors"
	"time"
	"auth/domain"
	"auth/storage"
)

func VerifyCode(entryToken, code string) (domain.AuthStatus, error) {
	codeState, ok := storage.GetCode(code)
	if !ok {
		deny(entryToken)
		return domain.StatusDenied, ErrCodeNotFound
	}

	if time.Now().After(codeState.ExpiresAt) {
		deny(entryToken)
		return domain.StatusDenied, ErrCodeExpired
	}

	if codeState.EntryToken != entryToken {
		deny(entryToken)
		return domain.StatusDenied, ErrInvalidCode
	}

	authState, ok := storage.GetAuthState(entryToken)
	if !ok {
		return domain.StatusDenied, errors.New("auth state not found")
	}

	authState.Status = domain.StatusApproved
	storage.SaveAuthState(authState, entryToken)

	return authState.Status, nil
}

func deny(entryToken string) {
	if authState, ok := storage.GetAuthState(entryToken); ok {
		authState.Status = domain.StatusDenied
		storage.SaveAuthState(authState, entryToken)
	}
}
