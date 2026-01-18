package clients

import (
	"encoding/json"
	"net/http"
	"time"
)

type GithubEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

// Primary email — это главный email аккаунта GitHub.
// Если primary == false, это вторичный email, его нельзя использовать как основной идентификатор пользователя
// Verified email означает: пользователь подтвердил, что владеет этим email’ом

func GetGithubEmails(token string) ([]GithubEmail, error) {
	// Не http.Get, потому что нам нужно управлять заголовками и передавать токен
	// Создаем объект запроса
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user/emails",
		nil,
	)
	if err != nil {
		return nil, err
	}


	req.Header.Set("Authorization", "Bearer " + token) // OAuth2 стандарт. Без этого заголовка будет 401
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()


	var emails []GithubEmail
	if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		return nil, err
	}

	return emails, nil
}
