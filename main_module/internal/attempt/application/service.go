package application

import (
	attemptDomain "main_module/internal/attempt/domain"
)

type AttemptService struct {
	Repo attemptDomain.Repository
}

func NewAttemptService(repo attemptDomain.Repository) *AttemptService {
	return &AttemptService{
		Repo: repo,
	}
}