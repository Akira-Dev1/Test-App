#pragma once
#include <string>

struct Course {
    int id;
    std::string title;
    std::string description;
    int author_id;
    bool is_deleted;
};
