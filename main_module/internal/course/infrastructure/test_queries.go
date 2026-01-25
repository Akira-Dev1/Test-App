package infrastructure

import (
	"errors"

	courseDomain 	"main_module/internal/course/domain"
)

// Получение информации о курсе (Список тестов)
func (r *PostgresRepo) GetAllTests(courseID string) ([]courseDomain.Test, error) {
	rows, err := r.DB.Query(`
		SELECT id, title 
		FROM tests 
		WHERE course_id = $1 AND is_deleted = false`,
		courseID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []courseDomain.Test
	for rows.Next() {
		var t courseDomain.Test
		rows.Scan(&t.ID, &t.Title)
		tests = append(tests, t)
	}

	return tests, nil
}

// Получение информации о тесте (Активный тест или нет)
func (r *PostgresRepo) GetStatus(testID string) (courseDomain.Test, error) {
	var test courseDomain.Test
	err := r.DB.QueryRow(`
		SELECT is_active 
		FROM tests 
		WHERE id = $1 AND is_deleted = false`,
		testID,
	).Scan(&test.IsActive)
	if err != nil {
		return courseDomain.Test{}, err
	}

	return test, nil
}

// Изменение теста (Активировать/Деактивировать тест)
func (r *PostgresRepo) UpdateStatus(testID string, status bool) error {
	res, err := r.DB.Exec(`
		UPDATE tests 
		SET is_active = $1 
		WHERE id = $2 AND is_deleted = false`,
		status,
		testID,
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

// Создание теста
func (r *PostgresRepo) CreateNewTest(courseID string, title string) (courseDomain.Test, error) {
	var test courseDomain.Test
	err := r.DB.QueryRow(`
        INSERT INTO tests (course_id, title)
        VALUES ($1, $2) 
		RETURNING id`,
		courseID,
		title,
	).Scan(&test.ID)

	if err != nil {
		return courseDomain.Test{}, err
	}

	return test, err
}

// Удаление теста
func (r *PostgresRepo) DeleteTestByID(testID string) error {
	res, err := r.DB.Exec(`
		UPDATE tests 
		SET is_deleted = true 
		WHERE id = $1`,
		testID,
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
