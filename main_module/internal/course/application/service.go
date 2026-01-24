package application

import (
	"errors"
	"slices"

	courseDomain "main_module/internal/course/domain"
	authDomain "main_module/internal/auth/domain"
)

type AccessRule struct {
	Permission		string 	// разрешение

	AllowSelf		bool 	// для себя
	AllowOwner		bool 	// если автор
	AllowIfEnrolled	bool 	// если записан

	DefaultAllow		bool	 // по умолчанию
}

type CourseService struct {
	Repo courseDomain.Repository
}

func NewCourseService(repo courseDomain.Repository) *CourseService {
	return &CourseService{
		Repo: repo,
	}
}
func (s *CourseService) GetCourses(user *authDomain.UserContext) ([]courseDomain.Course, error) {
	rule := AccessRule{
		DefaultAllow: true,
	}

	if !CheckAccess(user, rule, "") {
		return nil, errors.New("forbidden")
	}

	return s.Repo.GetAll()
}

func CheckAccess(user *authDomain.UserContext, rule AccessRule, ownerID string) bool {
	if user.Blocked {
		return false
	}

	if rule.Permission != "" && slices.Contains(user.Permissions, rule.Permission) {
		return true
	}

	if rule.AllowOwner && user.UserID == ownerID {
		return true
	}

	return rule.DefaultAllow
}