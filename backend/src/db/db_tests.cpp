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
    if (success && !isActive) {
        finalizeAllTestAttempts(testId);
    }
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



