package services

import (
	"auth/clients"
	"auth/domain"
	"auth/permissions"
	"auth/storage"
	"errors"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GithubAuth(code string) (*domain.User, error) {
    token, err := clients.ExchangeCodeForToken(code)
    if err != nil {
        return nil, err
    }

    profile, err := clients.GetGithubProfile(token)
    if err != nil {
        return nil, err
    }
	// Если данные пришли - достаем ВСЕ почты пользователя (может быть > 1) и ищем нужную 
    emails, err := clients.GetGithubEmails(token)
    if err != nil {
        return nil, err
    }

    var email string
    for _, e := range emails {
        if e.Primary && e.Verified {
            email = e.Email
            break
        }
    }

    if email == "" {
        return nil, errors.New("no verified primary email")
    }

    user, err := storage.FindUserByEmail(email)
    if err == nil {
        // пользователь уже есть
        return user, nil
    }

    if err != mongo.ErrNoDocuments {
        // реальная ошибка БД
        log.Printf("FindUserByEmail error: %v\n", err)

        return nil, err
    }

    // пользователя нет — создаём
    githubID := strconv.FormatInt(profile.ID, 10) // 10 - десятичная система

    newUser, err := storage.CreateUser(domain.User{
		Email:         email,
		GithubID:      &githubID,
		// Roles:         []domain.Role{domain.RoleStudent},
        Roles:         []string{string(domain.RoleStudent)},
        Permissions: permissions.ResolvePermissions([]string{string(domain.RoleStudent)}),
		RefreshTokens: []string{},
		CreatedAt:     time.Now(),
    })
    if err != nil {
        return nil, err
    }

    return newUser, nil
}

