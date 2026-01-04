#include "db.h"



// Получение всех тестов
std::vector<Test> DB::getTests() {
    ensureConnection();
    std::vector<Test> tests;
    PGresult* res = PQexec((PGconn*)conn, "SELECT id, title, description FROM tests");

    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        std::cerr << "SELECT failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
        PQclear(res);
        return tests;
    }
    for (int i = 0; i < PQntuples(res); i++) {
        const char* id_val = PQgetvalue(res, i, 0);
        const char* name_val = PQgetvalue(res, i, 1);
        const char* desc_val = PQgetvalue(res, i, 2);
        tests.push_back({
            id_val ? std::stoi(id_val) : 0,
            name_val ? name_val : "",
            desc_val ? desc_val : ""
        });
    }
    PQclear(res);
    return tests;
}


// Получение теста по айди
Test DB::getTestById(int testId) {
    ensureConnection();
    std::string idStr = std::to_string(testId);
    const char* paramValues[1] = { idStr.c_str() };

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "SELECT id, title, description FROM tests WHERE id = $1",
        1, nullptr, paramValues, nullptr, nullptr, 0
    );

    if (PQresultStatus(res) != PGRES_TUPLES_OK || PQntuples(res) == 0) {
        PQclear(res);
        throw std::runtime_error("Test not found");
    }

    Test t{
        std::stoi(PQgetvalue(res, 0, 0)),
        PQgetvalue(res, 0, 1),
        PQgetvalue(res, 0, 2)
    };
    PQclear(res);
    return t;
}


// Получение тестов по айди курса
std::vector<Test> DB::getTestsByCourse(int courseId) {
    ensureConnection();
    std::vector<Test> tests;
    std::string idStr = std::to_string(courseId);
    const char* paramValues[1];
    paramValues[0] = idStr.c_str();

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "SELECT id, title, description FROM tests WHERE course_id = $1",
        1, nullptr,
        paramValues,
        nullptr, nullptr, 0
    );

    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        std::cerr << "SELECT tests by course failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
        PQclear(res);
        return tests;
    }
    
    for (int i = 0; i < PQntuples(res); i++) {
        tests.push_back({
            std::stoi(PQgetvalue(res, i, 0)),
            PQgetvalue(res, i, 1),
            PQgetvalue(res, i, 2)
        });
    }

    PQclear(res);
    return tests;
}


// Создание теста (привязанного к курсу)
void DB::createTest(int courseId, const std::string& title, const std::string& description, int authorId) {
    ensureConnection();
    const char* paramValues[4] = { std::to_string(courseId).c_str(), title.c_str(), description.c_str(), std::to_string(authorId).c_str() };

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "INSERT INTO tests(course_id, title, description, author_id) VALUES ($1, $2, $3, $4)",
        3, nullptr, paramValues, nullptr, nullptr, 0
    );

    if (PQresultStatus(res) != PGRES_COMMAND_OK) {
        std::cerr << "Create test failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
    }
    PQclear(res);
}


// Удаление теста
void DB::deleteTest(int testId) {
    ensureConnection();
    std::string idStr = std::to_string(testId);
    const char* paramValues[1] = { idStr.c_str() };

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "UPDATE tests SET is_deleted = true WHERE id = $1",
        1, nullptr, paramValues, nullptr, nullptr, 0
    );

    if (PQresultStatus(res) != PGRES_COMMAND_OK) {
        std::cerr << "DELETE test failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
    }
    PQclear(res);
}

// Установка активности теста
void DB::setTestActivity(int testId, bool active) {
    ensureConnection();
    std::string idStr = std::to_string(testId);
    std::string activeStr = active ? "true" : "false";
    const char* paramValues[2] = { idStr.c_str(), activeStr.c_str() };

    const char* sql = 
        "BEGIN; "
        "UPDATE tests SET is_active = $2 WHERE id = $1; "
        "UPDATE test_attempts SET status = 'completed' "
        "WHERE test_id = $1 AND status = 'in_progress' AND $2 = false; "
        "COMMIT;";

    PGresult* res = PQexecParams((PGconn*)conn, sql, 2, nullptr, paramValues, nullptr, nullptr, 0);

    if (PQresultStatus(res) != PGRES_COMMAND_OK) {
        std::cerr << "Set activity failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
    }
    PQclear(res);
}