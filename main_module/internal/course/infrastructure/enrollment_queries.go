package infrastructure

import (
	"errors"
	courseDomain 	"main_module/internal/course/domain"
)

// Получение информации о курсе (Список студентов)
func (r *PostgresRepo) GetStudents(courseID string) (courseDomain.StudentsIDS, error) {
	rows, err := r.DB.Query(`
        SELECT cs.user_id 
		FROM course_students cs
        JOIN courses c ON cs.course_id = c.id
        WHERE c.id = $1 AND c.is_deleted = false`,
		courseID,
	)
	if err != nil {
		return courseDomain.StudentsIDS{}, err
	}
	defer rows.Close()

	var students courseDomain.StudentsIDS
	for rows.Next() {
		var s string
		rows.Scan(&s)

		students.IDS = append(students.IDS, s)
	}

	return students, nil
}

// Запись пользователя на курс
func (r *PostgresRepo) AddUserToCourse(courseID string, targetUserID string) error {
	var exists bool
    err := r.DB.QueryRow(`
		SELECT EXISTS(
			SELECT 1 
			FROM courses 
			WHERE id = $1 AND is_deleted = false
		)`, 
		courseID,
		).Scan(&exists)

    if err != nil {
        return err
    }
    if !exists {
        return errors.New("course not found or deleted")
    }

    _, err = r.DB.Exec(`
        INSERT INTO course_students (course_id, user_id)
        VALUES ($1::int, $2)
        ON CONFLICT (course_id, user_id) DO NOTHING`,
        courseID, targetUserID,
    )

	return err
}

// Отчисление пользователя с курса
func (r *PostgresRepo) RemoveUserFromCourse(courseID string, targetUserID string) error {
	res, err := r.DB.Exec(`
		DELETE FROM course_students 
		WHERE course_id = $1::int AND user_id = $2`,
		courseID,
		targetUserID,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("record not found")
	}

	return nil
}

// Записан ли пользователь на курс
func (r *PostgresRepo) IsEnrolled(courseID string, targetUserID string) (bool, error) {
	var exists bool
	
	err := r.DB.QueryRow(`
		SELECT EXISTS(
			SELECT 1 
			FROM course_students 
			WHERE course_id = $1::int AND user_id = $2
		)`, 
		courseID, 
		targetUserID,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
