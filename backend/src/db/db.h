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
    Test getTestById(int testId);
    bool updateTestStatus(int testId, bool isActive);
    void finalizeAllTestAttempts(int testId);
    std::vector<Test> getTestsByCourseId(int courseId);
    int createTest(int courseId, const std::string& title, int authorId);
    bool deleteTest(int testId);
    void setTestActivity(int testId, bool active);
    bool isUserEnrolled(int courseId, int userId);

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
