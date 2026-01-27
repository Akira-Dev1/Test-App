package application

import (
	"errors"

	authDomain "main_module/internal/auth/domain"
	userDomain "main_module/internal/user/domain"
)

// Получить список пользователей
func (s *UserService) GetUserData(user *authDomain.UserContext, targetUserID string, hasCourses bool, hasTests bool, hasGrades bool) (userDomain.UserData, error) {
	rule := AccessRule{
		Permission: "user:data:read",
		AllowSelf: true,
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "") {
		return userDomain.UserData{}, errors.New("forbidden")
	}

	return s.Repo.GetData(targetUserID, hasCourses, hasTests, hasGrades)
}