package services

import (
	"errors"
	"strconv"

	"auth/clients"
	"auth/domain"
	"auth/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

func GithubAuth(code string) error {
	token, err := clients.ExchangeCodeForToken(code)
	if err != nil {
		return err
	}

	profile, err := clients.GetGithubProfile(token)
	if err != nil {
		return err
	}

	emails, err := clients.GetGithubEmails(token)
	if err != nil {
		return err
	}

	var email string
	for _, e := range emails {
		if e.Primary && e.Verified {
			email = e.Email
			break
		}
	}

	if email == "" {
		return errors.New("no verified primary email")
	}

	user, err := storage.FindUserByEmail(email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			githubID := strconv.FormatInt(profile.ID, 10)

			_, err = storage.CreateUser(domain.User{
				Email:    email,
				GithubID: &githubID,
			})
			return err
		}
		return err
	}

	_ = user // пока просто подтверждаем, что пользователь найден
	return nil
}
