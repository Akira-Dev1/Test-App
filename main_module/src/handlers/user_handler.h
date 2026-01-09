#pragma once
#include "crow.h"
#include "../db/db.h"
#include "../security/jwt.h"
#include "../security/access.h"
#include "../security/auth_guard.h"


inline void registerUserRoutes(crow::SimpleApp& app, DB& db) {
    // Информация о пользователе (курсы, оценки, тесты)
    CROW_ROUTE(app, "/users/<int>/data").methods("GET"_method)
    ([&db](const crow::request& req, int targetUserId) {
        UserContext ctx;
        auto auth = authGuard(req, ctx);
        if (auth.code != 200) return auth;

        bool isOwner = (ctx.userId == targetUserId);

        PermissionRule readRule{
            "user:data:read", 
            false, 
            nullptr};
        bool hasAdminPermission = (checkAccess(ctx, readRule, 0).code == 200);

        if (!isOwner && !hasAdminPermission) {
            return crow::response(403, "Forbidden: You can only view your own data");
        }

        bool includeCourses = req.url_params.get("courses") != nullptr;
        bool includeTests = req.url_params.get("tests") != nullptr;
        bool includeGrades = req.url_params.get("grades") != nullptr;

        if (!includeCourses && !includeTests && !includeGrades) {
            includeCourses = includeTests = includeGrades = true; 
        }

        auto data = db.getUserDataProfile(targetUserId, includeCourses, includeTests, includeGrades);

        return crow::response(200, std::move(data));
    });
}