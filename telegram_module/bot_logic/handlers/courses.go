package handlers

import (
	"bot_logic/main_module"
	"bot_logic/storage"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"log"
)

func getToken(r *http.Request) (string, string) {
	chatID := r.URL.Query().Get("chat_id")
	if chatID == "" {
		return "", ""
	}
	userData, _ := storage.AuthRedis.HGetAll(context.Background(), chatID).Result()
	return userData["access_token"], chatID
}

// Получить курсы
func GetCoursesHandler (w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)

	courses, err := main_module.GetCourses(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Получить курс по айди
func GetCourseHandler (w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	course_id, _ := strconv.Atoi(r.URL.Query().Get("course_id"))


	course, err := main_module.GetCourseByID(token, course_id)
	if err != nil {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

// Создать курс
func CreateCourseHandler (w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)

	var body struct {
		Title			string `json:"title"`
		Description		string `json:"description"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	course_id, err := main_module.CreateCourse(token, body.Title, body.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"course_id": course_id})
}

// Изменение курса
func UpdateCourseHandler (w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	course_id, _ := strconv.Atoi(r.URL.Query().Get("course_id"))

	var body struct {
		Title			string `json:"title"`
		Description		string `json:"description"`
	}
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "Invalid JSON body", http.StatusBadRequest)
        return
    }

	err := main_module.UpdateCourse(token, course_id, body.Title, body.Description)
    if err != nil {
        log.Printf("Update error: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusNoContent)
}

// Удаление курса
func DeleteCourseHandler (w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	course_id, _ := strconv.Atoi(r.URL.Query().Get("course_id"))

	err := main_module.DeleteCourse(token, course_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Запись пользователя на курс
func JoinCourseHandler (w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	
	var body struct {
		CourseID int    `json:"course_id"`
		UserID   string `json:"user_id"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	err := main_module.JoinCourse(token, body.CourseID, body.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Удаление пользователя с курса
func LeaveCourseHandler (w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	
	var body struct {
		CourseID int    `json:"course_id"`
		UserID   string `json:"user_id"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	err := main_module.LeaveCourse(token, body.CourseID, body.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
