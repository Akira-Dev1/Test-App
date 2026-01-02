#pragma once
#include <string>
#include <vector>
#include <unordered_set>

struct UserContext {
    int userId;
    bool blocked;
    std::unordered_set<std::string> permissions;
    std::unordered_set<std::string> roles;
};