package infrastructure

import (
	"database/sql"
	
	"main_module/internal/course/domain"
)

type PostgresRepo struct {
	DB *sql.DB
}

func NewCourseRepository(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		DB: db,
	}
}

func (r *PostgresRepo) GetAll() ([]domain.Course, error) {
	rows, err := r.DB.Query(`
		SELECT id, title, description 
		FROM courses 
		WHERE is_deleted = false
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		var c domain.Course
		rows.Scan(&c.ID, &c.Title, &c.Description)
		courses = append(courses, c)
	}

	return courses, nil
}