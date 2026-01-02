package domain

import "time"

type CodeState struct {
	Code       string
	EntryToken string
	ExpiresAt  time.Time
}
