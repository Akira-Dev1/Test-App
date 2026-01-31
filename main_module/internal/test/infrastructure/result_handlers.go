package infrastructure

import (
	"database/sql"
	"strconv"
	"encoding/json"
	"math"
	"errors"

	authDomain "main_module/internal/auth/domain"
	testDomain "main_module/internal/test/domain"
)

// Получить список пользователей, проходивших тест
func (r *PostgresRepo) GetTestAttemptUsers(testID int) (testDomain.StudentsIDS, error) {
    query := `
        SELECT DISTINCT user_id
        FROM test_attempts 
        WHERE test_id = $1
        ORDER BY user_id`

    rows, err := r.DB.Query(query, testID)
    if err != nil {
        return testDomain.StudentsIDS{}, err
    }
    defer rows.Close()

    var userIDs testDomain.StudentsIDS
    for rows.Next() {
        var userID string
        if err := rows.Scan(&userID); err != nil {
            return testDomain.StudentsIDS{}, err
        }
        userIDs.IDS = append(userIDs.IDS, userID)
    }

    return userIDs, rows.Err()
}

// Получить оценки за тест в процентах (0-100)
func (r *PostgresRepo) GetTestScores(user *authDomain.UserContext, testID int, isAuthor bool) (testDomain.Scores, error) {
    var maxScore int
    err := r.DB.QueryRow(`
        SELECT COALESCE(array_length(question_ids, 1), 0) 
        FROM tests 
        WHERE id = $1 AND is_deleted = false`, 
        testID,
    ).Scan(&maxScore)
    
    if err != nil {
        if err == sql.ErrNoRows {
            return testDomain.Scores{}, errors.New("test not found")
        }
        return testDomain.Scores{}, err
    }
    
    var rows *sql.Rows
    if isAuthor {
        rows, err = r.DB.Query(`
            SELECT user_id, score
            FROM test_attempts 
            WHERE test_id = $1 
              AND status = 'completed'`, 
            testID,
        )
    } else {
        rows, err = r.DB.Query(`
            SELECT user_id, score
            FROM test_attempts 
            WHERE test_id = $1 
              AND user_id = $2
              AND status = 'completed'`, 
            testID, user.UserID,
        )
    }
    
    if err != nil {
        return testDomain.Scores{}, err
    }
    defer rows.Close()
    
    var scores testDomain.Scores
    for rows.Next() {
        var score testDomain.Score
        var userID string
        var scoreVal int
        
        err := rows.Scan(&userID, &scoreVal)
        if err != nil {
            return testDomain.Scores{}, err
        }
        
        var percentage float64 = 0
        if maxScore > 0 {
            percentage = float64(scoreVal) / float64(maxScore) * 100
        }
        
        percentage = math.Round(percentage*10) / 10
        
        score.UserID = &user.UserID
        score.Score = &percentage
        scores.Scores = append(scores.Scores, score)
    }
    
    return scores, rows.Err()
}

// Получить ответы пользователей на тес
func (r *PostgresRepo) GetTestAttempts(user *authDomain.UserContext, testID int, isAuthor bool) (testDomain.Attempts, error) {
    var query string
    var args []interface{}
    
    query = `
        SELECT 
            id as attempt_id,
            user_id, 
            user_answers
        FROM test_attempts 
        WHERE test_id = $1`
    args = append(args, testID)
    
    if !isAuthor {
        query += " AND user_id = $2"
        args = append(args, user.UserID)
    }
    
    rows, err := r.DB.Query(query, args...)
    if err != nil {
        return testDomain.Attempts{}, err
    }
    defer rows.Close()
    
    var attempts testDomain.Attempts
    for rows.Next() {
        var attempt testDomain.Attempt
        var userIDStr string
        var attemptID int
        var answersJSON []byte
        
        err := rows.Scan(&attemptID, &userIDStr, &answersJSON)
        if err != nil {
            return testDomain.Attempts{}, err
        }
    
        attempt.UserID = &userIDStr
        attempt.AttemptID = &attemptID
        
        var answersMap map[string]interface{}
        if err := json.Unmarshal(answersJSON, &answersMap); err != nil {
            return testDomain.Attempts{}, err
        }
        
        for questionIDStr, answerValue := range answersMap {
            questionID, err := strconv.Atoi(questionIDStr)
            if err != nil {
                continue
            }
            
            var answerVal *int
            
            switch v := answerValue.(type) {
            case float64:
                val := int(v)
                answerVal = &val
            case int:
                answerVal = &v
            case string:
                if num, err := strconv.Atoi(v); err == nil {
                    answerVal = &num
                } else {
                    val := -1
                    answerVal = &val
                }
            default:
                answerVal = nil
            }
            
            if answerVal != nil {
                answer := testDomain.Answer{
                    QuestionID: &questionID,
                    Answer:     answerVal,
                }
                attempt.Answers = append(attempt.Answers, answer)
            }
        }
        
        attempts.Attempts = append(attempts.Attempts, attempt)
    }
    
    return attempts, rows.Err()
}