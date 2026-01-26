package application

import (
	"errors"

	authDomain "main_module/internal/auth/domain"
	courseDomain "main_module/internal/course/domain"
)

// Получить список курсов
func (s *CourseService) GetCourseStudents(user *authDomain.UserContext, courseID string) (courseDomain.StudentsIDS, error) {
	rule := AccessRule{
		Permission: "course:userList",
		AllowOwner: true,
		DefaultAllow: false,
	}

	course, err := s.Repo.GetByID(courseID)
	if err != nil {
		return courseDomain.StudentsIDS{}, errors.New("course not found")
	}

	if !CheckAccess(user, rule, *course.AuthorID, "") {
		return courseDomain.StudentsIDS{}, errors.New("forbidden")
	}

	return s.Repo.GetStudents(courseID)
}

// Запись пользователя на курс
func (s *CourseService) EnrollUser(user *authDomain.UserContext, courseID string, targetUserID string) error {
	rule := AccessRule{
		Permission: "course:user:add",
		AllowSelf: true,
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "", targetUserID) {
		return errors.New("forbidden")
	}

	return s.Repo.AddUserToCourse(courseID, targetUserID)
}

// Отчисление пользователя с курса
func (s *CourseService) UnenrollUser(user *authDomain.UserContext, courseID string, targetUserID string) error {
	rule := AccessRule{
		Permission: "course:user:add",
		AllowSelf: true,
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "", targetUserID) {
		return errors.New("forbidden")
	}

	return s.Repo.RemoveUserFromCourse(courseID, targetUserID)
}
