package application

import (
	questionDomain "main_module/internal/question/domain"
)

type QuestionService struct {
	Repo questionDomain.Repository
}

func NewQuestionService(repo questionDomain.Repository) *QuestionService {
	return &QuestionService{
		Repo: repo,
	}
}