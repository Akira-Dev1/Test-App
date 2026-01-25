package domain

import (
	authDomain "main_module/internal/auth/domain"
)

type Repository interface {
	GetAll() 	([]map[string]any, error)
	GetByID(id string) 	(map[string]any, error)
	UpdateByID(id string, title any, description any) error
	CreateNew(user *authDomain.UserContext, title string, description string) (map[string]string, error)
	DeleteByID(id string) error
}