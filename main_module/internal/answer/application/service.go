package application

import (
	answerDomain "main_module/internal/answer/domain"
)

type AnswerService struct {
	Repo answerDomain.Repository
}

func NewAnswerService(repo answerDomain.Repository) *AnswerService {
	return &AnswerService{
		Repo: repo,
	}
}