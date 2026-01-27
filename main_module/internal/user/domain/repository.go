package domain

import (
)

type Repository interface {
	// Profile
	GetName(targetUserID string) (UserData, error)
	UpdateName(targetUserID string, name string) error

	// UserData
	GetData(targetUserID string, hasCourses bool, hasTests bool, hasGrades bool) (UserData, error)

	// AdminPanel
	GetUsersIDS() (UsersIDS, error)
	GetRoles(targetUserID string) (UserData, error)
	UpdateRoles(targetUserID string, roles []string) error
	GetStatus(targetUserID string) (UserData, error)
	UpdateStatus(targetUserID string, status bool) error		

}