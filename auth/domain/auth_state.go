package domain

import "time"

type AuthStatus string

const (
	StatusPending  AuthStatus = "pending"
	StatusApproved AuthStatus = "approved"
	StatusDenied   AuthStatus = "denied"
)

type AuthState struct {
	EntryToken string
	ExpiresAt  time.Time
	Status     AuthStatus
	UserID        string // заполняется при success
	AccessToken   string // JWT access (позже)
	RefreshToken  string // JWT refresh (позже)
}
