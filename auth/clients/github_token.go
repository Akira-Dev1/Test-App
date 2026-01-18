package clients

import (
	// "bytes" //bytes – стандартная библиотека Go, нужен для bytes.NewBuffer, чтобы передать JSON как тело POST запроса.
	"encoding/json"
	"net/http"
	"os" // стандартная библиотека для чтения переменных окружения
	"time"
	"net/url"
	"strings"
	"errors"
	"io"
	"log"
)

type githubTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func ExchangeGithubCodeForToken(code string) (string, error) {
	data := url.Values{}
	data.Set("client_id", os.Getenv("GITHUB_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("GITHUB_CLIENT_SECRET"))
	data.Set("code", code)
	data.Set("redirect_uri", "https://elda-unsimple-leandro.ngrok-free.dev/login/github/callback")

	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Println("GitHub token raw response:", string(bodyBytes))

	var tokenResp githubTokenResponse

	if err := json.Unmarshal(bodyBytes, &tokenResp); err != nil {
		return "", err
	}

	if tokenResp.AccessToken == "" {
		return "", errors.New("github did not return access token")
	}

	return tokenResp.AccessToken, nil
}

// Bearer-токен означает: любой, у кого есть этот токен, считается авторизованным.
// То есть: нет подписи, нет шифрования, просто строка

// Зачем bytes.NewBuffer?
// HTTP-запрос в Go принимает тело как io.Reader, а не []byte.
// bytes.Buffer: 
// оборачивает []byte
// делает из него поток (io.Reader)