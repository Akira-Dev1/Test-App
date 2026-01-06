#include "db.h"

// Получение теста по айди ccc
Test DB::getTestById(int testId) {
    ensureConnection();
    
    std::string idStr = std::to_string(testId);
    const char* params[] = { idStr.c_str() };

    const char* sql = 
        "SELECT id, course_id, title, is_active, is_deleted FROM tests WHERE id = $1";
    PGresult* res = PQexecParams(
        (PGconn*)conn,
        sql,
        1, nullptr, params, nullptr, nullptr, 0
    );

    if (PQresultStatus(res) != PGRES_TUPLES_OK || PQntuples(res) == 0) {
        PQclear(res);
        return {0, 0, "", false, false}; 
    }

    Test t;
    t.id = std::stoi(PQgetvalue(res, 0, 0));
    t.course_id = std::stoi(PQgetvalue(res, 0, 1));
    t.title = PQgetvalue(res, 0, 2);
    t.is_active = (std::string(PQgetvalue(res, 0, 3)) == "t");
    t.is_deleted = (std::string(PQgetvalue(res, 0, 4)) == "t");

    PQclear(res);
    return t;
}


// Получение тестов по айди курса
std::vector<Test> DB::getTestsByCourseId(int courseId) {
    ensureConnection();

    std::string cId = std::to_string(courseId);
    const char* params[] = { cId.c_str() };

    const char* sql = 
        "SELECT id, title FROM tests WHERE course_id = $1 AND is_deleted = false";

    PGresult* res = PQexecParams(
        (PGconn*)conn, 
        sql,
        1, nullptr, params, nullptr, nullptr, 0
    );

    std::vector<Test> tests;

    if (PQresultStatus(res) == PGRES_TUPLES_OK) {
        for (int i = 0; i < PQntuples(res); i++) {
            tests.push_back(Test{
                std::stoi(PQgetvalue(res, i, 0)),
                courseId,
                PQgetvalue(res, i, 1),
                false,
                false
            });
        }
    }

    PQclear(res);
    return tests;
}


// Создание теста (привязанного к курсу) ccc
int DB::createTest(int courseId, const std::string& title, int authorId) {
    ensureConnection();

    std::string cIdStr = std::to_string(courseId);
    std::string aIdStr = std::to_string(authorId);
    const char* paramValues[3] = { 
        cIdStr.c_str(), 
        title.c_str(), 
        aIdStr.c_str() 
    };

    const char* sql = 
        "INSERT INTO tests (course_id, title, author_id) "
        "VALUES ($1, $2, $3) RETURNING id";
    
    PGresult* res = PQexecParams(
        (PGconn*)conn, 
        sql, 
        3, nullptr, paramValues, nullptr, nullptr, 0);

    int newId = -1;
    if (PQresultStatus(res) == PGRES_TUPLES_OK) {
        newId = std::stoi(PQgetvalue(res, 0, 0));
    } else {
        std::cerr << "Create test failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
    }
    PQclear(res);
    return newId;
}


// Удаление теста ccc
bool DB::deleteTest(int testId) {
    ensureConnection();

    std::string tIdStr = std::to_string(testId);
    const char* paramValues[] = { tIdStr.c_str() };

    const char* sql = 
        "UPDATE tests SET is_deleted = true WHERE id = $1";

    PGresult* res = PQexecParams(
        (PGconn*)conn, 
        sql, 
        1, nullptr, paramValues, nullptr, nullptr, 0
    );

    bool success = (PQresultStatus(res) == PGRES_COMMAND_OK && std::string(PQcmdTuples(res)) == "1");
    PQclear(res);
    return success;
}

// Установка активности теста ccc
bool DB::updateTestStatus(int testId, bool isActive) {
    ensureConnection();
    std::string tId = std::to_string(testId);
    const char* status = isActive ? "true" : "false";
    const char* params[] = { status, tId.c_str() };

    const char* sql = 
        "UPDATE tests SET is_active = $1 WHERE id = $2 AND is_deleted = false";
    PGresult* res = PQexecParams(
        (PGconn*)conn,
        sql,
        2, nullptr, params, nullptr, nullptr, 0
    );

    bool success = (PQresultStatus(res) == PGRES_COMMAND_OK && std::string(PQcmdTuples(res)) == "1");
    PQclear(res);
    return success;
}
// Завершение всех попыток
void DB::finalizeAllTestAttempts(int testId) {
    ensureConnection();
    std::string tId = std::to_string(testId);
    const char* params[] = { tId.c_str() };

    const char* sql = 
        "UPDATE test_attempts SET status = 'completed' "
        "WHERE test_id = $1 AND status = 'in_progress'";
    PQexecParams(
        (PGconn*)conn,
        sql,
        1, nullptr, params, nullptr, nullptr, 0
    );
}

// Проверка записи на курс ccc
bool DB::isUserEnrolled(int courseId, int userId) {
    ensureConnection();

    std::string cId = std::to_string(courseId);
    std::string uId = std::to_string(userId);
    const char* params[] = { cId.c_str(), uId.c_str() };

    const char* sql = 
        "SELECT 1 FROM course_students WHERE course_id = $1 AND user_id = $2";
    PGresult* res = PQexecParams(
        (PGconn*)conn, 
        sql,
        2, nullptr, params, nullptr, nullptr, 0
    );
    bool enrolled = (PQresultStatus(res) == PGRES_TUPLES_OK && PQntuples(res) > 0);
    PQclear(res);
    return enrolled;
}

// Удаление вопроса из теста
bool DB::removeQuestionFromTest(int testId, int questionId) {
    ensureConnection();
    std::string tId = std::to_string(testId);
    std::string qId = std::to_string(questionId);
    const char* tParams[] = { tId.c_str() };

    const char* checkSql = "SELECT 1 FROM test_attempts WHERE test_id = $1::int LIMIT 1";
    PGresult* checkRes = PQexecParams((PGconn*)conn, checkSql, 1, nullptr, tParams, nullptr, nullptr, 0);
    
    if (PQresultStatus(checkRes) != PGRES_TUPLES_OK) {
        PQclear(checkRes);
        return false;
    }
    
    bool hasAttempts = (PQntuples(checkRes) > 0);
    PQclear(checkRes);

    if (hasAttempts) {
        return false;
    }

    const char* sql = 
        "UPDATE tests "
        "SET question_ids = array_remove(question_ids, $2::int) "
        "WHERE id = $1::int";
    
    const char* params[] = { tId.c_str(), qId.c_str() };
    PGresult* res = PQexecParams((PGconn*)conn, sql, 2, nullptr, params, nullptr, nullptr, 0);
    
    bool success = (PQresultStatus(res) == PGRES_COMMAND_OK);
    PQclear(res);
    return success;
}


// Добавление вопроса в тест
bool DB::addQuestionToTest(int testId, int questionId) {
    ensureConnection();
    
    std::string tId = std::to_string(testId);
    std::string qId = std::to_string(questionId);
    const char* params[] = { tId.c_str() };

    const char* checkSql = "SELECT 1 FROM test_attempts WHERE test_id = $1::int LIMIT 1";
    PGresult* checkRes = PQexecParams((PGconn*)conn, checkSql, 1, nullptr, params, nullptr, nullptr, 0);

    if (PQresultStatus(checkRes) != PGRES_TUPLES_OK) {
        std::cerr << "SQL Error (check): " << PQerrorMessage((PGconn*)conn) << std::endl;
        PQclear(checkRes);
        return false; 
    }

    int rows = PQntuples(checkRes);
    PQclear(checkRes);
    
    if (rows > 0) {
        return false;
    }

    const char* sql = 
        "UPDATE tests "
        "SET question_ids = array_append(question_ids, $2::int) "
        "WHERE id = $1::int AND NOT ($2::int = ANY(question_ids))";
        
    const char* updateParams[] = { tId.c_str(), qId.c_str() };
    PGresult* res = PQexecParams((PGconn*)conn, sql, 2, nullptr, updateParams, nullptr, nullptr, 0);

    bool success = (PQresultStatus(res) == PGRES_COMMAND_OK);
    
    PQclear(res);
    return success;
}


// Изменение порядка вопросов в тесте
bool DB::reorderQuestionsInTest(int testId, const std::vector<int>& questionIds) {
    ensureConnection();
    
    std::string tId = std::to_string(testId);
    const char* tParams[] = { tId.c_str() };

    const char* checkSql = "SELECT 1 FROM test_attempts WHERE test_id = $1::int LIMIT 1";
    PGresult* checkRes = PQexecParams((PGconn*)conn, checkSql, 1, nullptr, tParams, nullptr, nullptr, 0);
    
    if (PQresultStatus(checkRes) != PGRES_TUPLES_OK) {
        PQclear(checkRes);
        return false;
    }
    
    bool hasAttempts = (PQntuples(checkRes) > 0);
    PQclear(checkRes);
    
    if (hasAttempts) return false;

    std::string arrayStr = "{";
    for (size_t i = 0; i < questionIds.size(); ++i) {
        arrayStr += std::to_string(questionIds[i]);
        if (i < questionIds.size() - 1) arrayStr += ",";
    }
    arrayStr += "}";

    const char* sql = "UPDATE tests SET question_ids = $1::int[] WHERE id = $2::int";
    const char* params[] = { arrayStr.c_str(), tId.c_str() };
    
    PGresult* res = PQexecParams((PGconn*)conn, sql, 2, nullptr, params, nullptr, nullptr, 0);
    
    bool success = (PQresultStatus(res) == PGRES_COMMAND_OK);
    if (!success) {
        std::cerr << "Reorder failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
    }

    PQclear(res);
    return success;
}


// Список пользователей прошедших тест
std::vector<int> DB::getUsersWhoPassedTest(int testId) {
    ensureConnection();
    std::string tId = std::to_string(testId);
    const char* params[] = { tId.c_str() };

    const char* sql = 
        "SELECT DISTINCT user_id FROM test_attempts "
        "WHERE test_id = $1::int AND status = 'completed'";

    PGresult* res = PQexecParams((PGconn*)conn, sql, 1, nullptr, params, nullptr, nullptr, 0);
    
    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        std::cerr << "Get passed users failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
        PQclear(res);
        return {};
    }

    std::vector<int> userIds;
    int rows = PQntuples(res);
    userIds.reserve(rows);

    for (int i = 0; i < rows; i++) {
        userIds.push_back(std::stoi(PQgetvalue(res, i, 0)));
    }

    PQclear(res);
    return userIds;
}


// Получить оценку пользователей (или себя)
std::vector<UserScore> DB::getTestScores(int testId, int userIdFilter, bool isAuthor) {
    ensureConnection();
    std::string tId = std::to_string(testId);
    
    std::string sql = "SELECT user_id, score FROM test_attempts WHERE test_id = $1::int AND status = 'completed'";
    
    int paramCount = 1;
    std::vector<const char*> params = { tId.c_str() };
    std::string uId;

    if (!isAuthor) {
        sql += " AND user_id = $2::int";
        uId = std::to_string(userIdFilter);
        params.push_back(uId.c_str());
        paramCount = 2;
    }

    PGresult* res = PQexecParams((PGconn*)conn, sql.c_str(), paramCount, nullptr, params.data(), nullptr, nullptr, 0);
    std::vector<UserScore> scores;

    if (PQresultStatus(res) == PGRES_TUPLES_OK) {
        for (int i = 0; i < PQntuples(res); i++) {
            scores.push_back({
                std::stoi(PQgetvalue(res, i, 0)),
                std::stod(PQgetvalue(res, i, 1))
            });
        }
    }
    PQclear(res);
    return scores;
}

// Посмотреть ответы пользователей (пользователя)
std::vector<AttemptDetails> DB::getTestAttemptDetails(int testId, int userIdFilter, bool isAuthor) {
    ensureConnection();
    
    std::string sql = 
        "SELECT ta.user_id, q.content, q.options->>(ans.selected_option_idx) as answer_text "
        "FROM test_attempts ta "
        "JOIN test_answers ans ON ta.id = ans.attempt_id "
        "JOIN questions q ON ans.question_id = q.id AND ans.question_version = q.version "
        "WHERE ta.test_id = $1::int";

    std::vector<const char*> params;
    std::string tId = std::to_string(testId);
    params.push_back(tId.c_str());
    std::string uId;

    if (!isAuthor) {
        sql += " AND ta.user_id = $2::int";
        uId = std::to_string(userIdFilter);
        params.push_back(uId.c_str());
    }

    sql += " ORDER BY ta.user_id, ta.id";

    PGresult* res = PQexecParams((PGconn*)conn, sql.c_str(), params.size(), nullptr, params.data(), nullptr, nullptr, 0);
    
    std::map<int, AttemptDetails> attemptMap;
    if (PQresultStatus(res) == PGRES_TUPLES_OK) {
        for (int i = 0; i < PQntuples(res); i++) {
            int uid = std::stoi(PQgetvalue(res, i, 0));
            std::string qText = PQgetvalue(res, i, 1);
            std::string aText = PQgetvalue(res, i, 2);

            attemptMap[uid].user_id = uid;
            attemptMap[uid].answers.push_back({qText, aText});
        }
    }
    PQclear(res);

    std::vector<AttemptDetails> result;
    for (auto const& [id, details] : attemptMap) result.push_back(details);
    return result;
}
