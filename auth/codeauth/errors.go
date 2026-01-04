package codeauth

import "errors"

var (
	ErrCodeNotFound = errors.New("code not found")
	ErrCodeExpired  = errors.New("code expired")
	ErrInvalidCode  = errors.New("invalid code")
)
