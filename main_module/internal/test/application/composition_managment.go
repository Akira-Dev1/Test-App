package application

import (
	"errors"

	authDomain 		"main_module/internal/auth/domain"
	testDomain 		"main_module/internal/test/domain"


)

// Удалить вопрос из теста
func (s *TestService) RemoveQuestionFromTest(user *authDomain.UserContext, testID int, questionID int) error {
	rule := AccessRule{
		Permission: "user:data:read",
		AllowOwnerCourse: true,
		DefaultAllow: false,
	}

	hasAttempts, err := s.Repo.TestHasAttempts(testID)
	if err != nil {
		return err
	}
	if hasAttempts {
		return errors.New("cannot modify test with existing attempts")
	}

    authorID, err := s.Repo.GetTestCourseAuthor(testID)
    if err != nil {
        return err
    }

    if !CheckAccess(user, rule, *authorID, "", false) {
        return errors.New("forbidden")
    }

    err = s.Repo.RemoveQuestion(testID, questionID)
    if err != nil {
        return err
    }

    return nil
}    
// Добавить вопрос в тест
func (s *TestService) AddQuestionToTest(user *authDomain.UserContext, testID int, questionID int) error {
    rule := AccessRule{
        Permission:        "test:quest:add",
		AllowOwnerCourseAndQuestion: true,
        DefaultAllow:      false,
    }

    hasAttempts, err := s.Repo.TestHasAttempts(testID)
    if err != nil {
        return err
    }
    if hasAttempts {
        return errors.New("cannot modify test with existing attempts")
    }

    authorID, err := s.Repo.GetTestCourseAuthor(testID)
    if err != nil {
        return err
    }

    questionAuthorID, err := s.Repo.GetQuestionAuthorID(questionID)
    if err != nil {
        return err
    }

    if !CheckAccess(user, rule, *authorID, *questionAuthorID, false) {
        return errors.New("forbidden")
    }

    err = s.Repo.AddQuestion(testID, questionID)
    if err != nil {
        return err
    }

    return nil
}
// Изменить порядок следования вопросов в тесте
func (s *TestService) ChangeOrderOfQuestions(user *authDomain.UserContext, testID int, questionIDs testDomain.TestQuestionIDS) error {
    rule := AccessRule{
        Permission:        "test:quest:update",
        AllowOwnerCourse:  true,
        DefaultAllow:      false,
    }

    hasAttempts, err := s.Repo.TestHasAttempts(testID)
    if err != nil {
        return err
    }
    if hasAttempts {
        return errors.New("cannot modify test with existing attempts")
    }

    testAuthorID, err := s.Repo.GetTestCourseAuthor(testID)
    if err != nil {
        return err
    }

    if !CheckAccess(user, rule, *testAuthorID, "", false) {
        return errors.New("forbidden")
    }

    existingQuestions, err := s.Repo.GetTestQuestions(testID)
    if err != nil {
        return err
    }

    existingMap := make(map[int]bool)
    for _, q := range existingQuestions.IDS {
        existingMap[q] = true
    }

    for _, qID := range questionIDs.IDS {
        if !existingMap[qID] {
            return errors.New("there is an extra identifier in your request")
        }
    }

    err = s.Repo.UpdateQuestionOrder(testID, questionIDs)
    if err != nil {
        return err
    }

    return nil
}
// Получить список вопросов в тесте
func (s *TestService) GetTestQuestions(user *authDomain.UserContext, testID int) (testDomain.TestQuestionIDS, error) {
    rule := AccessRule{
        AllowOwnerCourse:  true,
		AllowHasActiveAttempt: true,
        DefaultAllow:      false,
    }

	hasActiveAttempt, err := s.Repo.HasActiveTestAttempt(user, testID)
	if err != nil {
		return testDomain.TestQuestionIDS{}, err
	}

    testAuthorID, err := s.Repo.GetTestCourseAuthor(testID)
    if err != nil {
        return testDomain.TestQuestionIDS{}, err
    }

    if !CheckAccess(user, rule, *testAuthorID, "", hasActiveAttempt) {
        return testDomain.TestQuestionIDS{}, errors.New("forbidden")
    }

    return s.Repo.GetTestQuestions(testID)
}
