package infrastructure

import (
	"database/sql"
	"errors"
	courseDomain 	"main_module/internal/course/domain"
	authDomain 		"main_module/internal/auth/domain"
)

//Create(course domain.Course)
//Delete(id string) — (установка флага is_deleted = true)

// Получить список курсов
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
		var c courseDomain.Course
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

// Получить курс по айди
func (r *PostgresRepo) GetByID(id string) (map[string]any, error) {
	var course courseDomain.Course
	
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

// Изменить курс
func (r *PostgresRepo) UpdateByID(id string, title any, description any) error {
	res, err := r.DB.Exec(`
		UPDATE courses
		SET 
			title = COALESCE($1, title), 
			description = COALESCE($2, description)
		WHERE id = $3 AND is_deleted = false`,
		title,
		description,
		id,
	)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("course not found")
	}

	return nil
}

// Создать курс
func (r *PostgresRepo) CreateNew(user *authDomain.UserContext, title string, description string) (map[string]string, error) {
	var id string
	err := r.DB.QueryRow(`
		INSERT INTO courses(title, description, author_id) 
		VALUES ($1, $2, $3) 
		RETURNING id`,
		title,
		description,
		user.UserID,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return map[string]string{"id": id}, err
}

// Удалить курс
func (r *PostgresRepo) DeleteByID(id string) error {
	res, err := r.DB.Exec(`
		UPDATE courses 
		SET is_deleted = true 
		WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("course not found")
	}

	return nil
}