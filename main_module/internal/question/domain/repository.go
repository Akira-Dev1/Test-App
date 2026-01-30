package domain

import (
	authDomain "main_module/internal/auth/domain"
)

type Repository interface {
	GetAllQuestions() ([]Question, error)
	GetQuestionDetails(questionID int, version int) (Question, error)
	GetQuestionDetailsMaxVersion(questionID int) (Question, error)
	HasAttempt(user *authDomain.UserContext, questionID int) bool
	UpdateByID(
		questionID int, 
		title string, 
		content string, 
		options []string, 
		correctOption int,
	) error
	CreateNew(
		user *authDomain.UserContext,
		title string,
		content string,
		options []string,
		correctOption int,
	) (Question, error)
	DeleteByID(questionID int) error
}