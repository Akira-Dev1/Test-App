package infrastructure

import (
	"errors"
	"database/sql"
	"encoding/json"
	"fmt"

	authDomain "main_module/internal/auth/domain"
	questionDomain 	"main_module/internal/question/domain"
)

// Получить список вопросов
func (r *PostgresRepo) GetAllQuestions() ([]questionDomain.Question, error) {
    rows, err := r.DB.Query(`
		SELECT DISTINCT ON (id) id, version, author_id, title
		FROM questions 
		WHERE is_deleted = false
		ORDER BY id, version DESC`,
    )
    if err != nil {
        return nil, errors.New("query failed")
    }

    defer rows.Close()

	var questions []questionDomain.Question

	for rows.Next() {
		var question questionDomain.Question
		if err := rows.Scan(
			&question.ID, 
			&question.Version, 
			&question.AuthorID, 
			&question.Title,
		); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}
// Получить детали вопроса по версии
func (r *PostgresRepo) GetQuestionDetails(questionID int, version int) (questionDomain.Question, error) {
    var question questionDomain.Question
    var optionsJSON []byte 

	err := r.DB.QueryRow(`
		SELECT title, content, options, correct_option
		FROM questions
		WHERE id = $1 AND version = $2 AND is_deleted = false`,
		questionID,
		version,
	).Scan(
		&question.Title, 
		&question.Content, 
		&optionsJSON, 
		&question.CorrectAnswerOption,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return questionDomain.Question{}, errors.New("queston not found")
		}
		return questionDomain.Question{}, err
	}

    if err := json.Unmarshal(optionsJSON, &question.AnswerOptions); err != nil {
        return questionDomain.Question{}, fmt.Errorf("failed to parse options: %w", err)
    }

    return question, nil
}
// Получить детали вопроса максимальной версии (Вспомогательная функция)
func (r *PostgresRepo) GetQuestionDetailsMaxVersion(questionID int) (questionDomain.Question, error) {
	var question questionDomain.Question

	err := r.DB.QueryRow(`
		SELECT id, version, author_id, title, content, correct_option
		FROM questions
		WHERE id = $1 AND is_deleted = false
		ORDER BY version DESC 
        LIMIT 1`,
		questionID,
	).Scan(
		&question.ID, 
		&question.Version, 
		&question.AuthorID,
		&question.Title,
		&question.Content,
		&question.CorrectAnswerOption,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return questionDomain.Question{}, errors.New("queston not found")
		}
		return questionDomain.Question{}, err
	}

	return question, nil
}
// Была ли попытка пройти тест с данным вопросом (Вспомогательная функция)
func (r *PostgresRepo) HasAttempt(user *authDomain.UserContext, questionID int) bool {
    var exists int
    err := r.DB.QueryRow(`
        SELECT 1 
        FROM test_attempts ta
        JOIN tests t ON ta.test_id = t.id
        WHERE ta.user_id = $1 
        AND $2 = ANY(t.question_ids)
        LIMIT 1`, 
		user.UserID, 
		questionID,
	).Scan(&exists)

    if err != nil {
        if err == sql.ErrNoRows {
            return false
        }
        return false
    }

	return true
}
// Изменить вопрос
func (r *PostgresRepo) UpdateByID(
	questionID int, 
	title string, 
	content string, 
	options []string, 
	correctOption int,
) error {
    optionsJSON, err := json.Marshal(options)
    if err != nil {
        return err
    }

    _, err = r.DB.Exec(`
        WITH current_data AS (
            SELECT 
                COALESCE(MAX(version), 0) as max_version,
                COALESCE(
                    (SELECT author_id FROM questions WHERE id = $1 LIMIT 1),
                    'unknown'
                ) as author
            FROM questions 
            WHERE id = $1
        )
        INSERT INTO questions (id, version, author_id, title, content, options, correct_option, is_deleted)
        SELECT $1, max_version + 1, author, $2, $3, $4, $5, false
        FROM current_data`,
        questionID, 
		title, 
		content, 
		string(optionsJSON), 
		correctOption,
    )
  
    return nil
}
// Создать вопрос
func (r *PostgresRepo) CreateNew(
	user *authDomain.UserContext,
	title string,
	content string,
	options []string,
	correctOption int,
) (questionDomain.Question, error) {
    optionsJSON, err := json.Marshal(options)
    if err != nil {
        return questionDomain.Question{}, err
    }

    var questionData questionDomain.Question
    err = r.DB.QueryRow(`
        INSERT INTO questions (id, version, author_id, title, content, options, correct_option) 
        VALUES (nextval('questions_id_seq'), 1, $1, $2, $3, $4, $5) 
        RETURNING id
    `, user.UserID, 
	title, 
	content, 
	string(optionsJSON), 
	correctOption,
	).Scan(&questionData.ID)

    if err != nil {
        return questionDomain.Question{}, fmt.Errorf("failed to create question: %w", err)
    }
    
    return questionData, nil
}
// Удалить вопрос
func (r *PostgresRepo) DeleteByID(questionID int) error {
    var isUsed bool
    err := r.DB.QueryRow(`
		SELECT EXISTS(
            SELECT 1 FROM tests WHERE $1 = ANY(question_ids)
        )
    `, questionID).Scan(&isUsed)
    
    if err != nil {
        return err
    }
    
    if isUsed {
        return errors.New("the question is used in tests")
    }

    result, err := r.DB.Exec(`
        UPDATE questions SET is_deleted = true WHERE id = $1
    `, questionID)
    
    if err != nil {
        return err
    }

	rowsAffected, _ := result.RowsAffected()
	if !(rowsAffected > 0) {
		return errors.New("error deleting question")
	}
    return nil
}