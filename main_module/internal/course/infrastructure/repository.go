package infrastructure

import (
	"database/sql"
)

type PostgresRepo struct {
	DB *sql.DB
}

func NewCourseRepository(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		DB: db,
	}
}