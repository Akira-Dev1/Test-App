#pragma once
#include <vector>
#include <string>
#include "../domain/test.h"
#include "../domain/course.h"

class DB {
public:
    DB(const std::string& conninfo);
    ~DB();

    std::vector<Test> getTests();
    Test getTestById(int id);
    std::vector<Test> getTestsByCourse(int courseId);
    void createTest(int courseId, const std::string& name, const std::string& description);
    void deleteTest(int id);

    std::vector<Course> getCourses();
    Course getCourseById(int id);
    void createCourse(const std::string& name, const std::string& description);
    void deleteCourse(int id);

private:
    void ensureConnection();
    void* conn;
    std::string conninfo;   
};
