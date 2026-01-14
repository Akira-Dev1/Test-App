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
    token, err := clients.ExchangeGithubCodeForToken(code)
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

    
    githubID := strconv.FormatInt(profile.ID, 10) // 10 - десятичная система

    user, err := storage.FindUserByEmail(email)
    if err == nil {
		if user.GithubID == nil || *user.GithubID == "" {
			user.GithubID = &githubID
			err := storage.UpdateUserGithubID(user.ID, githubID)
			if err != nil {
				log.Printf("Failed to update user with GithubID: %v\n", err)
				// Продолжаем с существующим пользователем
			}
		} else if *user.GithubID != githubID {
			// GithubID уже есть, но не совпадает
			log.Printf("GithubID mismatch for user %s. Stored: %s, new: %s", 
				user.Email, *user.GithubID, githubID)
		}
		
		return user, nil
    }

    if err != mongo.ErrNoDocuments {
        // реальная ошибка БД
        log.Printf("FindUserByEmail error: %v\n", err)

        return nil, err
    }


    newUser, err := storage.CreateUser(domain.User{
		Email:             email,
		GithubID:          &githubID,
        Name:              "Anonymous" + strconv.FormatInt(profile.ID + 14, 10),
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

