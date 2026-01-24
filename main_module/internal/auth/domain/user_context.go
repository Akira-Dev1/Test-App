package domain

type UserContext struct {
    UserID      string      `json:"user_id"`
    Permissions []string `json:"permissions"`
    Blocked     bool     `json:"blocked"`
}