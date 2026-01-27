package application

import (
	"errors"

	authDomain "main_module/internal/auth/domain"
	userDomain "main_module/internal/user/domain"
)

// Получить список пользователей
func (s *UserService) GetUsers(user *authDomain.UserContext) (userDomain.UsersIDS, error) {
	rule := AccessRule{
		Permission: "user:list:read",
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "") {
		return userDomain.UsersIDS{}, errors.New("forbidden")
	}

	return s.Repo.GetUsersIDS()
}

// Получить информацию о пользователе (Список ролей)
func (s *UserService) GetUserRoles(user *authDomain.UserContext, targetUserID string) (userDomain.UserData, error) {
	rule := AccessRule{
		Permission: "user:roles:read",
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "") {
		return userDomain.UserData{}, errors.New("forbidden")
	}

	return s.Repo.GetRoles(targetUserID)
}

// Изменить информацию о пользователе (Список ролей)
func (s *UserService) UpdateUserRoles(user *authDomain.UserContext, targetUserID string, roles []string) error {
	rule := AccessRule{
		Permission: "user:roles:write",
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "") {
		return errors.New("forbidden")
	}

	return s.Repo.UpdateRoles(targetUserID, roles)
}

// Получить информацию о пользователе (Заблокирован ли пользователь)
func (s *UserService) GetUserStatus(user *authDomain.UserContext, targetUserID string) (userDomain.UserData, error) {
	rule := AccessRule{
		Permission: "user:block:read",
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "") {
		return userDomain.UserData{}, errors.New("forbidden")
	}

	return s.Repo.GetStatus(targetUserID)
}

// Изменить информацию о пользователе (Заблокировать/Разблокировать пользователя)
func (s *UserService) UpdateUserStatus(user *authDomain.UserContext, targetUserID string, status bool) error {
	rule := AccessRule{
		Permission: "user:block:write",
		DefaultAllow: false,
	}

	if !CheckAccess(user, rule, "") {
		return errors.New("forbidden")
	}

	return s.Repo.UpdateStatus(targetUserID, status)
}