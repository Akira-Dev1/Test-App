#pragma once
#include "user_context.h"
#include <string>

// Проверка прав

struct PermissionRule {
    std::string permission;
    bool defaultAllowed = false;
    std::function<bool(const UserContext&, int)> checkDefault;
};

inline bool checkAccess(
    const UserContext& ctx,
    const PermissionRule& rule,
    int resourceOwnerId = -1
) {
    if (ctx.blocked) {
        return false;
    }
    if (ctx.permissions.find(rule.permission) != ctx.permissions.end()) {
        return true;
    }

    if (rule.defaultAllowed) {
        if (!rule.checkDefault) {
            return true;
        }
        return rule.checkDefault(ctx, resourceOwnerId);
    }
    return false;
}

inline bool hasPermission(
    const UserContext& ctx,
    const std::string& permission,
    int resourceOwnerId = -1
) {
    PermissionRule rule;
    rule.permission = permission;
    rule.defaultAllowed = false;
    return checkAccess(ctx, rule, resourceOwnerId);
}