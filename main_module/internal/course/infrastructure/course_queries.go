package infrastructure

import (
	"database/sql"
	"errors"
	courseDomain 	"main_module/internal/course/domain"
	authDomain 		"main_module/internal/auth/domain"
)

// Получить список курсов
func (r *PostgresRepo) GetAll() ([]courseDomain.Course, error) {
	rows, err := r.DB.Query(`
		SELECT id, title, description 
		FROM courses 
		WHERE is_deleted = false
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []courseDomain.Course
	for rows.Next() {
		var c courseDomain.Course
		rows.Scan(&c.ID, &c.Title, &c.Description)

		courses = append(courses, c)
	}

	return courses, nil
}

// Получить курс по айди
func (r *PostgresRepo) GetByID(id string) (courseDomain.Course, error) {
	var course courseDomain.Course
	
	err := r.DB.QueryRow(`
		SELECT title, description, author_id
		FROM courses 
		WHERE id = $1 AND is_deleted = false`,
		id,
	).Scan(&course.Title, &course.Description, &course.AuthorID)

	if err != nil {
		if err == sql.ErrNoRows {
			return courseDomain.Course{}, errors.New("course not found")
		}
		return courseDomain.Course{}, err
	}

	return course, nil
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
func (r *PostgresRepo) CreateNew(user *authDomain.UserContext, title string, description string) (courseDomain.Course, error) {
	var course courseDomain.Course
	err := r.DB.QueryRow(`
		INSERT INTO courses(title, description, author_id) 
		VALUES ($1, $2, $3) 
		RETURNING id`,
		title,
		description,
		user.UserID,
	).Scan(&course.ID)

	if err != nil {
		return courseDomain.Course{}, err
	}

	return course, err
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