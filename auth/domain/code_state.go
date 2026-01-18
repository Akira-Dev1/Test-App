package domain

import "time"

type CodeState struct {
	EntryToken string
	ExpiresAt  time.Time
}


type VerifyCodeRequest struct {
    Code       string `json:"code"`
    RefreshToken string `json:"refresh_token"`
}

type VerifyCodeResponse struct {
    Status string `json:"status"` 
}
