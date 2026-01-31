package application

import (
	"errors"

	authDomain "main_module/internal/auth/domain"
	testDomain "main_module/internal/test/domain"
)

// Получить список пользователей прошедших тест
func (s *TestService) GetPassedUsers(user *authDomain.UserContext, testID int) (testDomain.StudentsIDS, error) {
    rule := AccessRule{
        Permission:        "test:answer:read",
        AllowOwnerCourse:  true,
        DefaultAllow:      false,
    }

    testAuthorID, err := s.Repo.GetTestCourseAuthor(testID)
    if err != nil {
        return testDomain.StudentsIDS{}, err
    }

    if !CheckAccess(user, rule, *testAuthorID, "", false) {
        return testDomain.StudentsIDS{}, errors.New("forbidden")
    }

    userIDs, err := s.Repo.GetTestAttemptUsers(testID)
    if err != nil {
        return testDomain.StudentsIDS{}, err
    }

    return userIDs, nil
}
// Получить детали вопроса
func (s *TestService) GetUserRatings(user *authDomain.UserContext, testID int) (testDomain.Scores, error) {
    rule := AccessRule{
        Permission:        "test:answer:read",
        AllowOwnerCourse:  true,
        DefaultAllow:      false,
    }

    testAuthorID, err := s.Repo.GetTestCourseAuthor(testID)
    if err != nil {
        return testDomain.Scores{}, err
    }

    isAuthor := CheckAccess(user, rule, *testAuthorID, "", false)

    grade, err := s.Repo.GetTestScores(user, testID, isAuthor)
    if err != nil {
        return testDomain.Scores{}, err
    }

    return grade, nil
}
// Получить ответы пользователя
func (s *TestService) GetUserAnswers(user *authDomain.UserContext, testID int) (testDomain.Attempts, error) {
    rule := AccessRule{
        Permission:        "test:answer:read",
        AllowOwnerCourse:  true,
        DefaultAllow:      false,
    }

    testAuthorID, err := s.Repo.GetTestCourseAuthor(testID)
    if err != nil {
        return testDomain.Attempts{}, err
    }

    isAuthor := CheckAccess(user, rule, *testAuthorID, "", false)

    attempts, err := s.Repo.GetTestAttempts(user, testID, isAuthor)
    if err != nil {
        return testDomain.Attempts{}, err
    }

    return attempts, nil
}
