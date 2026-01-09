#include "jwt.h"
#include <jwt-cpp/jwt.h>
#include <cstdlib>

// Реализация проверки JWT

UserContext parseAndVerifyJWT(const crow::request& req) {
    const char* env_secret = std::getenv("JWT_SECRET_KEY");

    if (!env_secret) {
        throw std::runtime_error("JWT_SECRET_KEY is not defined in environment");
    }
    
    std::string jwt_secret(env_secret);

    auto auth = req.get_header_value("Authorization");
    if (auth.rfind("Bearer ", 0) != 0) {
        throw std::runtime_error("No token");
    }

    std::string token = auth.substr(7);

    auto decoded = jwt::decode(token);

    jwt::verify()
        .allow_algorithm(jwt::algorithm::hs256{jwt_secret})
        .verify(decoded);
    
    UserContext ctx;
    ctx.userId = decoded.get_payload_claim("user_id").as_integer();
    ctx.blocked = decoded.get_payload_claim("blocked").as_boolean();

    for (auto& p : decoded.get_payload_claim("permissions").as_array()) {
        ctx.permissions.insert(p.get<std::string>());
    }
    for (auto& r : decoded.get_payload_claim("roles").as_array()) {
        ctx.roles.insert(r.get<std::string>());
    }

    return ctx;
}