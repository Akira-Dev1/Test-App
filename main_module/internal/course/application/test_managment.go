package application

import (
	"errors"

	authDomain "main_module/internal/auth/domain"
	courseDomain "main_module/internal/course/domain"
)

// Получение информации о курсе (Список тестов)
func (s *CourseService) GetCourseTests(user *authDomain.UserContext, courseID string) ([]courseDomain.Test, error) {
	rule := AccessRule{
		Permission: "course:testList",
		AllowOwner: true,
		AllowIfEnrolled: true,
		DefaultAllow: false,
	}

	course, err := s.Repo.GetByID(courseID)
	if err != nil {
		return nil, errors.New("course not found")
	}

	if !CheckAccess(user, rule, *course.AuthorID) {
		return nil, errors.New("forbidden")
	}

	return s.Repo.GetAllTests(courseID)
}

// Получение информации о тесте (Активный тест или нет)
func (s *CourseService) GetTestStatus(user *authDomain.UserContext, courseID string, testID string) (courseDomain.Test, error) {
	rule := AccessRule{
		Permission: "course:test:read",
		AllowOwner: true,
		AllowIfEnrolled: true,
		DefaultAllow: false,
	}

	course, err := s.Repo.GetByID(courseID)
	if err != nil {
		return courseDomain.Test{}, errors.New("course not found")
	}

	if !CheckAccess(user, rule, *course.AuthorID) {
		return courseDomain.Test{}, errors.New("forbidden")
	}

	return s.Repo.GetStatus(testID)
}

// Изменение теста (Активировать/Деактивировать тест)
func (s *CourseService) UpdateTestStatus(user *authDomain.UserContext, courseID string, testID string, status bool) error {
	rule := AccessRule{
		Permission: "course:test:write",
		AllowOwner: true,
		DefaultAllow: false,
	}

	course, err := s.Repo.GetByID(courseID)
	if err != nil {
		return errors.New("course not found")
	}

	if !CheckAccess(user, rule, *course.AuthorID) {
		return errors.New("forbidden")
	}

	return s.Repo.UpdateStatus(testID, status)
}

// Создание теста
func (s *CourseService) CreateTest(user *authDomain.UserContext, courseID string, title string) (courseDomain.Test, error) {
	rule := AccessRule{
		Permission: "course:test:add",
		AllowOwner: true,
		DefaultAllow: false,
	}

	course, err := s.Repo.GetByID(courseID)
	if err != nil {
		return courseDomain.Test{}, errors.New("course not found")
	}

	if !CheckAccess(user, rule, *course.AuthorID) {
		return courseDomain.Test{}, errors.New("forbidden")
	}

	return s.Repo.CreateNewTest(courseID, title)
}

// Удаление теста
func (s *CourseService) DeleteTest(user *authDomain.UserContext, courseID string, testID string) error {
	rule := AccessRule{
		Permission: "course:test:del",
		AllowOwner: true,
		DefaultAllow: false,
	}

	course, err := s.Repo.GetByID(courseID)
	if err != nil {
		return errors.New("course not found")
	}

	if !CheckAccess(user, rule, *course.AuthorID) {
		return errors.New("forbidden")
	}

	return s.Repo.DeleteTestByID(testID)
}