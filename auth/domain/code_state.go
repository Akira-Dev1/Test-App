package domain

import "time"

type CodeState struct {
	Code       string
	EntryToken string
	ExpiresAt  time.Time
}


type VerifyCodeRequest struct {
    EntryToken string `json:"entry_token"`
    Code       string `json:"code"`
}

type VerifyCodeResponse struct {
    Status string `json:"status"` // approved / denied
}
