#pragma once
#include <crow.h>
#include "jwt.h"

// Проверка JWT

inline crow::response authGuard(
    const crow::request& req,
    UserContext& ctx
) {
    try {
        ctx = parseAndVerifyJWT(req);

        if (ctx.blocked) {
            return crow::response(418);
        }

        return crow::response(200);
    } catch (...) {
        return crow::response(401);
    }
}