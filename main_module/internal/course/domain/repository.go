package domain

type Repository interface {
	GetAll() ([]Course, error)
}