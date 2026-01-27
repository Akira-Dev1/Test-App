package application

import (
	"errors"

	authDomain "main_module/internal/auth/domain"
	userDomain "main_module/internal/user/domain"
)

// Получить информацию о пользователе (ФИО)
func (s *UserService) GetUserName(user *authDomain.UserContext, targetUserID string) (userDomain.UserData, error) {
	rule := AccessRule{
		DefaultAllow: true,
	}

	if !CheckAccess(user, rule, "") {
		return userDomain.UserData{}, errors.New("forbidden")
	}

	return s.Repo.GetName(targetUserID)
}

// Изменить информацию о пользователе (ФИО)
func (s *UserService) UpdateUserName(user *authDomain.UserContext, targetUserID string, name string) error {
	rule := AccessRule{
		Permission: "user:fullName:write",
		AllowSelf: true,
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "") {
		return errors.New("forbidden")
	}

	return s.Repo.UpdateName(targetUserID, name)
}