#include "crow.h"
#include "db/db.h"
#include "handlers/test_handler.h"
#include "handlers/course_handler.h"
#include "handlers/question_handler.h"
#include "handlers/attempt_handler.h"

int main() {
    crow::SimpleApp app;

    DB db("host=postgres port=5432 dbname=core_db user=core password=core");


    // Проверка активации
    CROW_ROUTE(app, "/health")([] {
        return "OK";
    });
    registerCourseRoutes(app, db);
    registerTestRoutes(app, db);
    registerQuestionRoutes(app, db);
    registerAttemptRoutes(app, db);

    app.port(18080).multithreaded().run();
}
