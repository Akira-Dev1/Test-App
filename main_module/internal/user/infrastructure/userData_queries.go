package infrastructure

import (
	userDomain 	"main_module/internal/user/domain"
)

// Получить информацию о пользователе (курсы, оценки, тесты)
func  (r *PostgresRepo) GetData(targetUserID string, hasCourses bool, hasTests bool, hasGrades bool) (userDomain.UserData, error) {
	var userData userDomain.UserData
	if hasCourses {
		rowsCourses, err := r.DB.Query(`
			SELECT c.id, c.title, c.description
			FROM courses c
			JOIN course_students cs ON c.id = cs.course_id
			WHERE cs.user_id = $1 AND c.is_deleted = false`,
			targetUserID,
		)

		if err != nil {
			return userDomain.UserData{}, err
		}

		defer rowsCourses.Close()

		var courses []userDomain.Course
		for rowsCourses.Next() {
			var course userDomain.Course
			if err := rowsCourses.Scan(
				&course.CourseID, 
				&course.Title, 
				&course.Description,
			); err != nil {
				return userDomain.UserData{}, err
			}
			courses = append(courses, course)
		}
		userData.Courses = courses
	}

	if hasTests {
		rowsTests, err := r.DB.Query(`
			SELECT t.id, t.title, t.course_id
			FROM tests t
			JOIN course_students cs ON t.course_id = cs.course_id
			WHERE cs.user_id = $1 AND t.is_deleted = false AND t.is_active = true`,
			targetUserID,
		)

		if err != nil {
			return userDomain.UserData{}, err
		}

		defer rowsTests.Close()

		var tests []userDomain.Test
		for rowsTests.Next() {
			var	test userDomain.Test
			if err := rowsTests.Scan(
				&test.TestID, 
				&test.Title, 
				&test.CourseID,
			); err != nil {
				return userDomain.UserData{}, err
			}
			tests = append(tests, test)
		}
		userData.Tests = tests
	}

	if hasGrades {
		rowsGrades, err := r.DB.Query(`
			SELECT 
				c.title,
				t.title,
				ta.created_at,
				ta.status,
				ta.score,
				ta.max_score,
			FROM test_attempts ta
			JOIN tests t ON ta.test_id = t.id
			JOIN courses c ON t.course_id = c.id
			WHERE ta.user_id = $1 
			AND t.is_deleted = false
			AND c.is_deleted = false`,
			targetUserID,
		)

		if err != nil {
			return userDomain.UserData{}, err
		}

		defer rowsGrades.Close()

		var grades []userDomain.Grade
		for rowsGrades.Next() {
			var	grade userDomain.Grade
			if err := rowsGrades.Scan(
				&grade.CourseTitle, 
				&grade.TestTitle, 
				&grade.Date, 
				&grade.Status, 
				&grade.Score, 
				&grade.MaxScore,
			); err != nil {
				return userDomain.UserData{}, err
			}
			grades = append(grades, grade)
		}
		userData.Grades = grades
	}
	return userData, nil
}