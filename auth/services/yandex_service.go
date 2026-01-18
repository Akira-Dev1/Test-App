package services

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"auth/domain"
	"auth/permissions"
	"auth/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type YandexUserInfo struct {
	ID           string   `json:"id"`
	DefaultEmail string   `json:"default_email"`
	Emails       []string `json:"emails"`
}

func ExchangeYandexCodeForToken(code string) (string, error) {
	clientID := os.Getenv("YANDEX_CLIENT_ID")
	clientSecret := os.Getenv("YANDEX_CLIENT_SECRET")
	redirectURI := os.Getenv("YANDEX_REDIRECT_URI")
	
	if redirectURI == "" {
		redirectURI = "http://localhost:8081/auth/yandex/callback"
	}
	
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectURI)
	
	req, err := http.NewRequest(
		"POST",
		"https://oauth.yandex.ru/token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", err
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	bodyBytes, _ := io.ReadAll(resp.Body)
	
	var tokenResp struct {
		AccessToken string `json:"access_token"`
		Error       string `json:"error"`
	}
	
	if err := json.Unmarshal(bodyBytes, &tokenResp); err != nil {
		return "", err
	}
	
	if tokenResp.Error != "" {
		return "", errors.New("yandex token error: " + tokenResp.Error)
	}
	
	if tokenResp.AccessToken == "" {
		return "", errors.New("no access token from yandex")
	}
	
	return tokenResp.AccessToken, nil
}

// getYandexUserInfo получает только ID и email
func getYandexUserInfo(accessToken string) (*YandexUserInfo, error) {
	req, err := http.NewRequest("GET", "https://login.yandex.ru/info", nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Authorization", "OAuth "+accessToken)
	
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	bodyBytes, _ := io.ReadAll(resp.Body)
	
	var userInfo YandexUserInfo
	if err := json.Unmarshal(bodyBytes, &userInfo); err != nil {
		return nil, err
	}
	
	if userInfo.ID == "" {
		return nil, errors.New("no user id from yandex")
	}
	
	// Если default_email пустой, пробуем взять первый из emails
	if userInfo.DefaultEmail == "" && len(userInfo.Emails) > 0 {
		userInfo.DefaultEmail = userInfo.Emails[0]
	}
	
	log.Printf("Got Yandex user: ID=%s, Email=%s", userInfo.ID, userInfo.DefaultEmail)
	
	return &userInfo, nil
}

func YandexAuth(code string) (*domain.User, error) {
    token, _ := ExchangeYandexCodeForToken(code)

	userInfo, err := getYandexUserInfo(token)
	if err != nil {
		return nil, err
	}
	
	
    user, err := storage.FindUserByEmail(userInfo.DefaultEmail)
    if err == nil {
		if user.YandexID == nil || *user.GithubID == "" {
			user.YandexID = &userInfo.ID
			err := storage.UpdateUserYandexID(user.ID, userInfo.ID)
			if err != nil {
				log.Printf("Failed to update user with YandexID: %v\n", err)
			}
		} else if *user.YandexID != userInfo.ID {
			log.Printf("YandexID mismatch for user %s. Stored: %s, new: %s", 
				user.Email, *user.YandexID, userInfo.ID)
		}
		
		return user, nil
    }

    if err != mongo.ErrNoDocuments {
        // реальная ошибка БД
        log.Printf("FindUserByEmail error: %v\n", err)

        return nil, err
    }


    newUser, err := storage.CreateUser(domain.User{
		Email:             userInfo.DefaultEmail,
		YandexID:          &userInfo.ID,
        Name:              "Anonymous" + strconv.FormatInt(143948752, 10),
        Roles:             []string{string(domain.RoleStudent)},
        Permissions:       permissions.ResolvePermissions([]string{string(domain.RoleStudent)}),
		RefreshTokens:     []string{},
        IsBlocked:         false,
		CreatedAt:         time.Now(),
    })
    if err != nil {
        return nil, err
    }

    return newUser, nil
}