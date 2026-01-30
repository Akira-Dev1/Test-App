package application

import (
	"errors"

	authDomain "main_module/internal/auth/domain"
	questionDomain "main_module/internal/question/domain"
)

// Получить список вопросов
func (s *QuestionService) GetQuestions(user *authDomain.UserContext) ([]questionDomain.Question, error) {
	rule := AccessRule{
		Permission: "user:data:read",
		AllowOwner: true,
		DefaultAllow: false,
	}

	questions, err := s.Repo.GetAllQuestions()
	if err != nil {
		return nil, err
	}

	var resultQuestions []questionDomain.Question

	for _, question := range questions {
		if CheckAccess(user, rule, *question.AuthorID, false) {
			resultQuestions = append(resultQuestions, question)
		}
	}

	return resultQuestions, nil
}    
// Получить детали вопроса
func (s *QuestionService) GetQuestionByID(user *authDomain.UserContext, questionID int, version int) (questionDomain.Question, error) {
	rule := AccessRule{
		Permission: "quest:read",
		AllowHasAttempt: true,
		DefaultAllow: false,
	}

	question, err := s.Repo.GetQuestionDetailsMaxVersion(questionID)
    if err != nil {
        return questionDomain.Question{}, err
    }
	hasAttempt := s.Repo.HasAttempt(user, questionID)

	if !CheckAccess(user, rule, *question.AuthorID, hasAttempt) {
		return questionDomain.Question{}, errors.New("forbidden")
	}

	return s.Repo.GetQuestionDetails(questionID, version)
}
// Изменить вопрос
func (s *QuestionService) UpdateQuestion(
		user *authDomain.UserContext, 
		questionID int, 
		title string, hasTitle bool,
		content string, hasContent bool,
		options []string, hasOptions bool,
		correctOption int, hasCorrectOptions bool,
) error {
	rule := AccessRule{
		Permission: "quest:update",
		AllowOwner: true,
		DefaultAllow: false,
	}

	question, err := s.Repo.GetQuestionDetailsMaxVersion(questionID)
	if err != nil {
		return err
	}

	if !CheckAccess(user, rule, *question.AuthorID, false) {
		return errors.New("forbidden")
	}

	if !hasTitle {
		title = *question.Title
	}
	if !hasContent {
		content = *question.Content
	}
	if !hasOptions {
		options = question.AnswerOptions
	}
	if !hasCorrectOptions {
		correctOption = *question.CorrectAnswerOption
	}

	return s.Repo.UpdateByID(
		questionID,
		title,
		content,
		options,
		correctOption,
	)
}
// Создать вопрос
func (s *QuestionService) CreateQuestion(
		user *authDomain.UserContext, 
		title string,
		content string,
		options []string,
		correctOption int,
) (questionDomain.Question, error) {
	rule := AccessRule{
		Permission: "quest:create",
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "", false) {
		return questionDomain.Question{}, errors.New("forbidden")
	}

	return s.Repo.CreateNew(
		user,
		title,
		content,
		options,
		correctOption,
	)
}
// Удалить вопрос
func (s *QuestionService) DeleteQuestion(user *authDomain.UserContext, questionID int) error {
	rule := AccessRule{
		Permission: "quest:del",
		AllowOwner: true,
		DefaultAllow: false,
	}

	question, err := s.Repo.GetQuestionDetailsMaxVersion(questionID)
	if err != nil {
		return err
	}
	
	if !CheckAccess(user, rule, *question.AuthorID, false) {
		return errors.New("forbidden")
	}

	return s.Repo.DeleteByID(questionID)
}