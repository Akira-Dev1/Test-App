#pragma once
#include "crow.h"
#include "../db/db.h"
#include "../security/jwt.h"
#include "../security/access.h"
#include "../security/auth_guard.h"

inline void registerCourseRoutes(crow::SimpleApp& app, DB& db) {

    // Получение всех курсов (название, описание)
    CROW_ROUTE(app, "/courses").methods("GET"_method)
    ([&db](const crow::request& req){
        auto courses = db.getCourses();
        std::vector<crow::json::wvalue> course_list;
        for (auto& c : courses) {
            crow::json::wvalue item;
            item["id"] = c.id;
            item["name"] = c.name;
            item["description"] = c.description;
            course_list.push_back(std::move(item));
        }
        crow::json::wvalue final_res;
        final_res["courses"] = std::move(course_list);
        return final_res;
    });

    // Создание курса
    CROW_ROUTE(app, "/courses").methods("POST"_method)
    ([&db](const crow::request& req) {
        auto userIdPtr = req.url_params.get("current_user_id");
        if (!userIdPtr) {
            return crow::response(401, "User not identified");
        }
        int userId = std::stoi(userIdPtr);
        UserContext ctx;
        auto res = authGuard(req, ctx);
        if (res.code != 200) {
            return res;
        }
        PermissionRule rule {
            "course:add",
            false,
            [](const UserContext& ctx, int ownerId) {
                return ctx.userId == ownerId;
            }
        };
        if (!checkAccess(ctx, rule, userId)) {
            return res;
        } 
        auto body = crow::json::load(req.body);
        if (!body) {
            return crow::response(400, "Invalid JSON");
        }
        db.createCourse(body["name"].s(), body["description"].s());
        return crow::response(201);
    });

    // Удаление курса
    CROW_ROUTE(app, "/courses/<int>").methods("DELETE"_method)
    ([&db](const crow::request& req, int id) {
        auto userIdPtr = req.url_params.get("current_user_id");
        if (!userIdPtr) {
            return crow::response(401, "User not identified");
        }
        int userId = std::stoi(userIdPtr);
        UserContext ctx;
        auto res = authGuard(req, ctx);
        if (res.code != 200) {
            return res;
        }
        PermissionRule rule {
            "course:del",
            false,
            [](const UserContext& ctx, int ownerId) {
                return ctx.userId == ownerId;
            }
        };
        if (!checkAccess(ctx, rule, userId)) {
            return res;
        } 
        db.deleteTest(id);
        return crow::response(204);
    });
}
