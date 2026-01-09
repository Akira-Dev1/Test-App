#pragma once
#include "crow.h"
#include <vector>
#include <string>
#include <libpq-fe.h>
#include <stdexcept>
#include <iostream>
#include "../domain/test.h"
#include "../domain/course.h"
#include "../domain/question.h"

// Структура оценки пользователя
struct UserScore {
    int user_id;
    double score;
};

// Структура информации о пользователе
struct UserProfileData {
    std::vector<Course> courses;
    std::vector<Test> tests;
    std::vector<UserScore> grades;
};

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

    bool removeQuestionFromTest(int testId, int questionId);
    bool addQuestionToTest(int testId, int questionId);
    bool reorderQuestionsInTest(int testId, const std::vector<int>& questionIds);
    std::vector<int> getUsersWhoPassedTest(int testId);
    std::vector<UserScore> getTestScores(int testId, int userIdFilter, bool isAuthor);
    std::vector<AttemptDetails> getTestAttemptDetails(int testId, int userIdFilter, bool isAuthor);

    // Попытки
    int startTestAttempt(int testId, int userId);
    bool updateAttemptAnswer(int attemptId, int questionId, int answerIndex);
    bool isAttemptOwnedBy(int attemptId, int userId);
    bool completeAttempt(int attemptId);
    crow::json::wvalue getAttemptData(int testId, int userId);
    crow::json::wvalue getAttemptAnswers(int testId, int userId);

    // Курсы
    std::vector<Course> getCourses();
    Course getCourseById(int courseId);
    int createCourse(const std::string& title, const std::string& description, int teacherId);
    void deleteCourse(int courseId);
    bool updateCourse(int courseId, std::string title, std::string description);
    bool addStudentToCourse(int courseId, int userId);
    bool removeStudentFromCourse(int courseId, int userId);
    std::vector<int> getStudentIdsByCourseId(int courseId);
    bool isUserEnrolled(int courseId, int userId);

    // Вопросы
    int createQuestion(int authorId, const std::string& title, const std::string& content, const std::vector<std::string>& options, int correctOption);
    bool deleteQuestion(int questionId);
    int updateQuestion(int questionId, int userId, const std::string& title, const std::string& content, const std::vector<std::string>& options, int correctOption);
    std::vector<Question> getQuestionsList(int userId, bool canSeeAll);
    Question getQuestionById(int questionId);
    Question getQuestionByIdAndVersion(int questionId, int version);
    bool hasUserAttemptForQuestion(int userId, int questionId);

    // Пользователь
    crow::json::wvalue getUserDataProfile(int userId, bool includeCourses, bool includeTests, bool includeGrades);

private:
    void ensureConnection();
    void* conn;
    std::string conninfo;   
};
