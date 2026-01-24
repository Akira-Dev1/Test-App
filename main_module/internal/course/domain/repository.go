package domain

type Repository interface {
	GetAll() 	([]map[string]any, error)
	GetByID(id string) 	(map[string]any, error)
}