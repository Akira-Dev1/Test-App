package infrastructure

import (
	"database/sql"
	"errors"
	"strconv"

	authDomain "main_module/internal/auth/domain"
	testDomain "main_module/internal/test/domain"

	"github.com/lib/pq"
)

// Получить список вопросов
func (r *PostgresRepo) TestHasAttempts(testID int) (bool, error) {
    var hasAttempts bool
    err := r.DB.QueryRow(`
		SELECT EXISTS (
            SELECT 1 
            FROM test_attempts 
            WHERE test_id = $1
        )`, 
		testID,
	).Scan(&hasAttempts)
    
    if err != nil {
        return false, err
    }
    
    return hasAttempts, nil
}
// Удалить вопрос из теста
func (r *PostgresRepo) RemoveQuestion(testID int, questionID int) error {
    result, err := r.DB.Exec(`
        UPDATE tests 
        SET question_ids = array_remove(question_ids, $2)
        WHERE id = $1`, 
		testID, 
		questionID,
	)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return errors.New("test not found or question not in test")
    }
    
    return nil
}
// Добавить вопрос в тест
func (r *PostgresRepo) AddQuestion(testID int, questionID int) error {
    result, err := r.DB.Exec(`
        UPDATE tests 
        SET question_ids = array_append(question_ids, $2)
        WHERE id = $1 AND NOT ($2 = ANY(question_ids))`, 
		testID, 
		questionID,
	)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        exists, err := r.IsQuestionInTest(testID, questionID)
        if err == nil && exists {
            return errors.New("question already in test")
        }
        return errors.New("test not found")
    }
    
    return nil
}

// Проверить, есть ли вопрос в тесте (Вспомогательная)
func (r *PostgresRepo) IsQuestionInTest(testID int, questionID int) (bool, error) {
    var exists bool
    err := r.DB.QueryRow(`
        SELECT EXISTS (
            SELECT 1 FROM tests 
            WHERE id = $1 AND $2 = ANY(question_ids))`, 
		testID, 
		questionID,
	).Scan(&exists)
    
    if err != nil {
        return false, err
    }
    
    return exists, nil
}
// Изменить порядок вопросов в тесте
func (r *PostgresRepo) UpdateQuestionOrder(testID int, questionIDs testDomain.TestQuestionIDS) error {
    arrayLiteral := "{"
    for i, id := range questionIDs.IDS {
        arrayLiteral += strconv.Itoa(id)
        if i < len(questionIDs.IDS)-1 {
            arrayLiteral += ","
        }
    }
    arrayLiteral += "}"
    
    result, err := r.DB.Exec(`
        UPDATE tests 
        SET question_ids = $1
        WHERE id = $2`, 
		arrayLiteral, 
		testID,
	)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return errors.New("test not found")
    }
    
    return nil
}

func (r *PostgresRepo) GetTestQuestions(testID int) (testDomain.TestQuestionIDS, error) {
    var ids testDomain.TestQuestionIDS
	var pqArray pq.Int64Array
    err := r.DB.QueryRow(`
		SELECT question_ids FROM tests WHERE id = $1`, 
		testID,
	).Scan(&pqArray)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return testDomain.TestQuestionIDS{}, errors.New("test not found")
        }
        return testDomain.TestQuestionIDS{}, err
    }
    
	ids.IDS = make([]int, len(pqArray))
	for i, val := range pqArray {
		ids.IDS[i] = int(val)
	}

    return ids, nil
}
// Получить тест по ID
func (r *PostgresRepo) GetTestByID(testID int) (testDomain.Test, error) {
    var test testDomain.Test
    var questionIDs pq.Int64Array
    
    err := r.DB.QueryRow(`
        SELECT 
            id, 
            course_id, 
            title, 
            question_ids, 
            is_active, 
            is_deleted,
        FROM tests 
        WHERE id = $1 AND is_deleted = false`, 
		testID,
	).Scan(
        &test.ID,
        &test.CourseID,
        &test.Title,
        pq.Array(&questionIDs),
        &test.IsActive,
        &test.IsDeleted,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return testDomain.Test{}, errors.New("test not found")
        }
        return testDomain.Test{}, err
    }
    
	test.QuestionIDs = make([]int, len(questionIDs))
	for i, val := range questionIDs {
		test.QuestionIDs[i] = int(val)
	}
    
    return test, nil
}
// Получить только автора курса для теста
func (r *PostgresRepo) GetTestCourseAuthor(testID int) (*string, error) {
    query := `
        SELECT c.author_id
        FROM tests t
        JOIN courses c ON t.course_id = c.id
        WHERE t.id = $1 
          AND t.is_deleted = false
          AND c.is_deleted = false`

    var authorID string
    err := r.DB.QueryRow(query, testID).Scan(&authorID)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("test not found")
        }
        return nil, err
    }
    
    return &authorID, nil
}
// Получить автора вопроса по ID
func (r *PostgresRepo) GetQuestionAuthorID(questionID int) (*string, error) {
    var authorID string
    err := r.DB.QueryRow(`
        SELECT DISTINCT ON (id) author_id
        FROM questions 
        WHERE id = $1 
          AND is_deleted = false
        ORDER BY id, version DESC`, 
		questionID,
	).Scan(&authorID)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("question not found")
        }
        return nil, err
    }
    
    return &authorID, nil
}
// Проверить, есть ли у пользователя активная попытка теста
func (r *PostgresRepo) HasActiveTestAttempt(user *authDomain.UserContext, testID int) (bool, error) {
    var exists bool
    err := r.DB.QueryRow(`
        SELECT EXISTS (
            SELECT 1 
            FROM test_attempts 
            WHERE user_id = $1 
              AND test_id = $2 
              AND status = 'in_progress')`, 
		user.UserID, 
		testID,
	).Scan(&exists)
    
    if err != nil {
        return false, err
    }
    
    return exists, nil
}