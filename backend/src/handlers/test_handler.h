#pragma once
#include "crow.h"
#include "../db/db.h"

inline void registerTestRoutes(crow::SimpleApp& app, DB& db) {
    // Список всех тестов
    CROW_ROUTE(app, "/tests").methods("GET"_method)
    ([&db]{
        auto tests = db.getTests();
        std::vector<crow::json::wvalue> test_list;
        for (auto& t : tests) {
            crow::json::wvalue item;
            item["id"] = t.id;
            item["name"] = t.name;
            item["description"] = t.description;
            test_list.push_back(std::move(item));
        }
        crow::json::wvalue final_res;
        final_res["tests"] = std::move(test_list);
        return final_res;
    });


    // Получение тестов по курсу
    CROW_ROUTE(app, "/courses/<int>/tests").methods("GET"_method)
    ([&db](int courseId) {
        auto tests = db.getTestsByCourse(courseId);
        crow::json::wvalue res;
        std::vector<crow::json::wvalue> test_list;
        for (auto& t : tests) {
            crow::json::wvalue item;
            item["id"] = t.id;
            item["name"] = t.name;
            item["description"] = t.description;
            test_list.push_back(std::move(item));
        }
        crow::json::wvalue final_res;
        final_res["tests"] = std::move(test_list);
        return final_res;
    });
    // Создание теста по курсу
    CROW_ROUTE(app, "/courses/<int>/tests").methods("POST"_method)
    ([&db](const crow::request& req, int courseId) {
        auto body = crow::json::load(req.body);
        if (!body) {
            return crow::response(400, "Invalid JSON");
        }
        db.createTest(
            courseId,
            body["name"].s(),
            body["description"].s()
        );
        return crow::response(201);
    });


    // Удаление теста по id
    CROW_ROUTE(app, "/tests/<int>").methods("DELETE"_method)
    ([&db](int id) {
        db.deleteTest(id);
        return crow::response(204);
    });
}
