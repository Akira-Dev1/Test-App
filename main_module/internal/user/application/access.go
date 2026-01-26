package application

import (
	"slices"

	authDomain "main_module/internal/auth/domain"
)


type AccessRule struct {
	Permission		string 	// разрешение

	AllowSelf		bool 	// для себя
	AllowOwner		bool 	// если автор
	AllowIfEnrolled	bool 	// если записан

	DefaultAllow		bool	 // по умолчанию
}

func CheckAccess(user *authDomain.UserContext, rule AccessRule, ownerID string, targetUserID string) bool {
	if user.Blocked {
		return false
	}

	if rule.Permission != "" && slices.Contains(user.Permissions, rule.Permission) {
		return true
	}

	if rule.AllowSelf && user.UserID == targetUserID {
		return true
	}

	if rule.AllowOwner && user.UserID == ownerID {
		return true
	}

	return rule.DefaultAllow
}