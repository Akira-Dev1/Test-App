package application

import (
	"slices"

	authDomain "main_module/internal/auth/domain"
)


type AccessRule struct {
	Permission					string 	// разрешение

	AllowOwnerCourse			bool 	// если автор курса
	AllowOwnerCourseAndQuestion	bool 	// если автор курса и вопроса
	AllowHasActiveAttempt		bool	// если активная попытка

	DefaultAllow				bool	 // по умолчанию
}

func CheckAccess(user *authDomain.UserContext, rule AccessRule, ownerCourseID string, ownerQuestionID string, hasAttempt bool) bool {
	if user.Blocked {
		return false
	}

	if rule.Permission != "" && slices.Contains(user.Permissions, rule.Permission) {
		return true
	}

	if rule.AllowOwnerCourse && user.UserID == ownerCourseID {
		return true
	}
	if rule.AllowOwnerCourseAndQuestion && user.UserID == ownerCourseID && user.UserID == ownerQuestionID {
		return true
	}

    if rule.AllowHasActiveAttempt && hasAttempt {
        return true
    }

	return rule.DefaultAllow
}