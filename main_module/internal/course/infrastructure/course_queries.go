package infrastructure

import (
	"database/sql"
	"errors"
	"main_module/internal/course/domain"
)

//GetByID(id string)
//Create(course domain.Course)
//Update(course domain.Course)
//Delete(id string) — (установка флага is_deleted = true)

func (r *PostgresRepo) GetAll() ([]map[string]any, error) {
	rows, err := r.DB.Query(`
		SELECT id, title, description 
		FROM courses 
		WHERE is_deleted = false
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataCourses []map[string]any
	for rows.Next() {
		var c domain.Course
		rows.Scan(&c.ID, &c.Title, &c.Description)

		dataC := map[string]any {
			"id": 			c.ID,
			"title": 		c.Title,
			"description": 	c.Description,
		}

		dataCourses = append(dataCourses, dataC)
	}

	return dataCourses, nil
}

func (r *PostgresRepo) GetByID(id string) (map[string]any, error) {
	var course domain.Course
	
	err := r.DB.QueryRow(`
		SELECT title, description, author_id
		FROM courses 
		WHERE id = $1 AND is_deleted = false`,
		id,
	).Scan(&course.Title, &course.Description, &course.AuthorID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("course not found")
		}
		return nil, err
	}

	dataCourse := map[string]any {
		"title": 	course.Title,
		"description": 	course.Description,
		"author_id": 	course.AuthorID,
	}
	return dataCourse, nil
}