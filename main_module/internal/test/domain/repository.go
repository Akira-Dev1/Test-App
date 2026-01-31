package domain

import (
	authDomain "main_module/internal/auth/domain"
)

type Repository interface {
	// Composition
	TestHasAttempts(testID int) (bool, error)
	RemoveQuestion(testID int, questionID int) error
	AddQuestion(testID int, questionID int) error
	IsQuestionInTest(testID int, questionID int) (bool, error)
	UpdateQuestionOrder(testID int, questionIDs TestQuestionIDS) error
	GetTestQuestions(testID int) (TestQuestionIDS, error)
	GetTestByID(testID int) (Test, error)
	GetTestCourseAuthor(testID int) (*string, error)
	GetQuestionAuthorID(questionID int) (*string, error)
	HasActiveTestAttempt(user *authDomain.UserContext, testID int) (bool, error)

	// Results
	GetTestAttemptUsers(testID int) (StudentsIDS, error)
	GetTestScores(user *authDomain.UserContext, testID int, isAuthor bool) (Scores, error)
	GetTestAttempts(user *authDomain.UserContext, testID int, isAuthor bool) (Attempts, error)
}