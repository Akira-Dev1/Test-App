package application

import (
	"errors"

	authDomain "main_module/internal/auth/domain"
	courseDomain "main_module/internal/course/domain"
)

// Получить список курсов
func (s *CourseService) GetCourses(user *authDomain.UserContext) ([]courseDomain.Course, error) {
	rule := AccessRule{
		DefaultAllow: true,
	}

	if !CheckAccess(user, rule, "", "") {
		return nil, errors.New("forbidden")
	}

	return s.Repo.GetAll()
}

// Получить курс по айди
func (s *CourseService) GetCourseByID(user *authDomain.UserContext, id string) (courseDomain.Course, error) {
	rule := AccessRule{
		DefaultAllow: true,
	}

	if !CheckAccess(user, rule, "", "") {
		return courseDomain.Course{}, errors.New("forbidden")
	}

	return s.Repo.GetByID(id)
}

// Изменить курс
func (s *CourseService) UpdateCourse(user *authDomain.UserContext, id string, title any, description any) error {
	rule := AccessRule{
		Permission: "course:info:write",
		AllowOwner: true,
		DefaultAllow: false,
	}

	course, err := s.Repo.GetByID(id)
	if err != nil {
		return errors.New("course not found")
	}

	if !CheckAccess(user, rule, *course.AuthorID, "") {
		return errors.New("forbidden")
	}

	return s.Repo.UpdateByID(id, title, description)
}

// Создать курс
func (s *CourseService) CreateCourse(user *authDomain.UserContext, title string, description string) (courseDomain.Course, error) {
	rule := AccessRule{
		Permission: "course:add",
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "", "") {
		return courseDomain.Course{}, errors.New("forbidden")
	}

	return s.Repo.CreateNew(user, title, description)
}

// Удалить курс
func (s *CourseService) DeleteCourse(user *authDomain.UserContext, id string) error {
	rule := AccessRule{
		Permission: "course:del",
		AllowOwner: true,
		DefaultAllow: false,
	}

	course, err := s.Repo.GetByID(id)
	if err != nil {
		return errors.New("course not found")
	}

	if !CheckAccess(user, rule, *course.AuthorID, "") {
		return errors.New("forbidden")
	}

	return s.Repo.DeleteByID(id)
}
