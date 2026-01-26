package application

import (
	userDomain "main_module/internal/user/domain"
)

type UserService struct {
	Repo userDomain.Repository
}

func NewUserService(repo userDomain.Repository) *UserService {
	return &UserService{
		Repo: repo,
	}
}