#include "db.h"



// Добавить вопрос в конец теста
void DB::addQuestionToTest(int testId, int questionId) {
    ensureConnection();
    const char* paramValues[2] = { std::to_string(testId).c_str(), std::to_string(questionId).c_str() };

    PGresult* res = PQexecParams((PGconn*)conn,
        "UPDATE tests SET question_ids = array_append(question_ids, $2) "
        "WHERE id = $1 AND NOT EXISTS (SELECT 1 FROM test_attempts WHERE test_id = $1)",
        2, nullptr, paramValues, nullptr, nullptr, 0);

    if (PQresultStatus(res) != PGRES_COMMAND_OK || std::stoi(PQcmdTuples(res)) == 0) {
        std::cerr << "Add question failed (check if attempts already exist)" << std::endl;
    }
    PQclear(res);
}

// Изменить порядок вопросов 
void DB::updateQuestionsOrder(int testId, const std::string& pgArray) {
    ensureConnection();
    const char* paramValues[2] = { std::to_string(testId).c_str(), pgArray.c_str() };

    PGresult* res = PQexecParams((PGconn*)conn,
        "UPDATE tests SET question_ids = $2::integer[] "
        "WHERE id = $1 AND NOT EXISTS (SELECT 1 FROM test_attempts WHERE test_id = $1)",
        2, nullptr, paramValues, nullptr, nullptr, 0);

    PQclear(res);
}
