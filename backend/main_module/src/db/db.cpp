#include "db.h"
#include <libpq-fe.h>
#include <iostream>
#include <stdexcept>

// Конструктор
DB::DB(const std::string& conninfo) : conninfo(conninfo), conn(nullptr) {
    ensureConnection();
}
// Деструктор
DB::~DB() {
    if (conn) {
        PQfinish((PGconn*)conn);
    }
}
// Проверка подключения
void DB::ensureConnection() {
    if (!conn || PQstatus((PGconn*)conn) != CONNECTION_OK) {
        if (conn) PQfinish((PGconn*)conn);
        
        conn = PQconnectdb(conninfo.c_str());
        
        if (PQstatus((PGconn*)conn) != CONNECTION_OK) {
            std::cerr << "DB Connection Error: " << PQerrorMessage((PGconn*)conn) << std::endl;
        } else {
            std::cout << "Successfully connected to DB" << std::endl;
        }
    }
}
// Получение всех тестов
std::vector<Test> DB::getTests() {
    ensureConnection();
    std::vector<Test> tests;
    PGresult* res = PQexec((PGconn*)conn, "SELECT id, name, description FROM tests");

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
// Получение всех курсов
std::vector<Course> DB::getCourses() {
    ensureConnection();
    std::vector<Course> courses;
    PGresult* res = PQexec((PGconn*)conn, "SELECT id, name, description FROM courses");

    if (PQresultStatus(res) != PGRES_TUPLES_OK) {
        std::cerr << "SELECT failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
        PQclear(res);
        return courses;
    }

    for (int i = 0; i < PQntuples(res); i++) {
        courses.push_back({
            std::stoi(PQgetvalue(res, i, 0)),
            PQgetvalue(res, i, 1),
            PQgetvalue(res, i, 2)
        });
    }
    PQclear(res);
    return courses;
}


// Получение теста по айди из бд
Test DB::getTestById(int id) {
    ensureConnection();
    std::string idStr = std::to_string(id);
    const char* paramValues[1] = { idStr.c_str() };

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "SELECT id, name, description FROM tests WHERE id = $1",
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
// Получение курса по айди из бд
Course DB::getCourseById(int id) {
    ensureConnection();
    std::string idStr = std::to_string(id);
    const char* paramValues[1] = { idStr.c_str() };

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "SELECT id, name, description FROM courses WHERE id = $1",
        1, nullptr, paramValues, nullptr, nullptr, 0
    );

    if (PQresultStatus(res) != PGRES_TUPLES_OK || PQntuples(res) == 0) {
        PQclear(res);
        throw std::runtime_error("Course not found");
    }

    Course c{
        std::stoi(PQgetvalue(res, 0, 0)),
        PQgetvalue(res, 0, 1),
        PQgetvalue(res, 0, 2)
    };
    PQclear(res);
    return c;  
}
// Получение тестов по курсу из бд
std::vector<Test> DB::getTestsByCourse(int courseId) {
    ensureConnection();
    std::vector<Test> tests;
    std::string idStr = std::to_string(courseId);
    const char* paramValues[1];
    paramValues[0] = idStr.c_str();

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "SELECT id, name, description FROM tests WHERE course_id = $1",
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




// Создание теста в бд
void DB::createTest(int courseId, const std::string& name, const std::string& description) {
    ensureConnection();
    const char* paramValues[3] = { std::to_string(courseId).c_str(), name.c_str(), description.c_str() };

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "INSERT INTO tests(course_id, name, description) VALUES ($1, $2, $3)",
        3, nullptr, paramValues, nullptr, nullptr, 0
    );

    if (PQresultStatus(res) != PGRES_COMMAND_OK) {
        std::cerr << "Create test failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
    }
    PQclear(res);
}
// Создание курса в бд
void DB::createCourse(const std::string& name, const std::string& description) {
    ensureConnection();
    const char* paramValues[2] = { name.c_str(), description.c_str() };

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "INSERT INTO courses(name, description) VALUES ($1, $2)",
        2, nullptr, paramValues, nullptr, nullptr, 0
    );

    if (PQresultStatus(res) != PGRES_COMMAND_OK) {
        std::cerr << "Create course failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
    }
    PQclear(res);
}




// Удаление теста в бд
void DB::deleteTest(int id) {
    ensureConnection();
    std::string idStr = std::to_string(id);
    const char* paramValues[1] = { idStr.c_str() };

    PGresult* res = PQexecParams(
        (PGconn*)conn,
        "DELETE FROM courses WHERE id = $1",
        1, nullptr, paramValues, nullptr, nullptr, 0
    );

    if (PQresultStatus(res) != PGRES_COMMAND_OK) {
        std::cerr << "DELETE failed: " << PQerrorMessage((PGconn*)conn) << std::endl;
    }
    PQclear(res);
}
