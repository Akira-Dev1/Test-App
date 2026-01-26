package domain

import (
	authDomain "main_module/internal/auth/domain"
)

type Repository interface {
	// Course
	GetAll() ([]Course, error)
	GetByID(id string) (Course, error)
	UpdateByID(id string, title any, description any) error
	CreateNew(user *authDomain.UserContext, title string, description string) (Course, error)
	DeleteByID(id string) error

	// Test
	GetAllTests(courseID string) ([]Test, error)
	GetStatus(testID string) (Test, error)
	UpdateStatus(testID string, status bool) error
	CreateNewTest(courseID string, title string) (Test, error)
	DeleteTestByID(testID string) error

	// Enrollment
	GetStudents(courseID string) (StudentsIDS, error)
	AddUserToCourse(courseID string, targetUserID string) error
	RemoveUserFromCourse(courseID string, targetUserID string) error

	IsEnrolled(courseID string, targetUserID string) (bool, error)
}