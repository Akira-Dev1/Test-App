package application

import (
	"slices"

	authDomain "main_module/internal/auth/domain"
)


type AccessRule struct {
	Permission			string 	// разрешение

	AllowOwner			bool 	// если автор
	AllowHasAttempt		bool	// Если есть попытка ответа содержащая данный вопрос

	DefaultAllow		bool	// по умолчанию
}

func CheckAccess(user *authDomain.UserContext, rule AccessRule, ownerID string, hasAttempt bool) bool {
	if user.Blocked {
		return false
	}

	if rule.Permission != "" && slices.Contains(user.Permissions, rule.Permission) {
		return true
	}

	if rule.AllowOwner && user.UserID == ownerID {
		return true
	}

    if rule.AllowHasAttempt && hasAttempt {
        return true
    }

	return rule.DefaultAllow
}