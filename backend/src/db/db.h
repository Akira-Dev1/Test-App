#pragma once
#include <vector>
#include <string>
#include <libpq-fe.h>
#include <stdexcept>
#include <iostream>
#include "../domain/test.h"
#include "../domain/course.h"

class DB {
public:
    DB(const std::string& conninfo);
    ~DB();

    // Тетсты
    std::vector<Test> getTests();
    Test getTestById(int id);
    std::vector<Test> getTestsByCourse(int courseId);
    void createTest(int courseId, const std::string& name, const std::string& description, int authorId);
    void deleteTest(int id);
    void setTestActivity(int testId, bool active);

    // Курсы
    std::vector<Course> getCourses();
    Course getCourseById(int courseId);
    int createCourse(const std::string& title, const std::string& description, int teacherId);
    void deleteCourse(int courseId);
    bool updateCourse(int courseId, std::string title, std::string description);
    bool addStudentToCourse(int courseId, int userId);
    bool removeStudentFromCourse(int courseId, int userId);
    std::vector<int> getStudentIdsByCourseId(int courseId);

    // Вопросы
    void addQuestionToTest(int testId, int questionId);
    void updateQuestionsOrder(int testId, const std::string& pgArray);

    // Ответы

private:
    void ensureConnection();
    void* conn;
    std::string conninfo;   
};
