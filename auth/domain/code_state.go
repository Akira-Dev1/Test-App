package domain

import "time"

type CodeState struct {
	EntryToken string
	ExpiresAt  time.Time
}


type VerifyCodeRequest struct {
    // EntryToken string `json:"entry_token"`
    Code       string `json:"code"`
    RefreshToken string `json:"refresh_token"`
}

type VerifyCodeResponse struct {
    Status string `json:"status"` // approved / denied
}
