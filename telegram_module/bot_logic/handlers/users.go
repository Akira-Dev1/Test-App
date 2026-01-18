package handlers

import (
	"encoding/json"
	"net/http"
	"bot_logic/main_module"
)

// Получить данные о пользователе (курсы, тесты, оценки)
func GetUserDataProfileHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	targetID := r.URL.Query().Get("target_id")
	
	c := r.URL.Query().Get("courses") == "true"
	t := r.URL.Query().Get("tests") == "true"
	g := r.URL.Query().Get("grades") == "true"

	data, err := main_module.GetUserData(token, targetID, c, t, g)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// Получить список пользователей
func GetUsersListHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	list, err := main_module.GetUserList(token)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	json.NewEncoder(w).Encode(list)
}

// Получить ФИО пользователя
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	targetID := r.URL.Query().Get("target_id")

	if targetID == "" {
		http.Error(w, "target_id is required", http.StatusBadRequest)
		return
	}

	info, err := main_module.GetUserInfo(token, targetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// Изменить ФИО пользователя
func UpdateFullNameHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	targetID := r.URL.Query().Get("target_id")

	var body struct { NewName string `json:"full_name"` }
	json.NewDecoder(r.Body).Decode(&body)

	err := main_module.UpdateUserFullName(token, targetID, body.NewName)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
}

// Получить роли пользователя
func GetUserRolesHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	targetID := r.URL.Query().Get("target_id")

	roles, err := main_module.GetUserRoles(token, targetID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string][]string{"roles": roles})
}

// Изменить роли пользователя
func UpdateUserRolesHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	targetID := r.URL.Query().Get("target_id")

	var body struct {
		Roles []string `json:"roles"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := main_module.UpdateUserRoles(token, targetID, body.Roles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Посмотреть статус пользователя (Заблокирован/Разблокирован)
func GetUserBlockStatusHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	targetID := r.URL.Query().Get("target_id")

	isBlocked, err := main_module.GetUserBlockStatus(token, targetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]bool{"is_blocked": isBlocked})
}

// Заблокировать/Разблокировать пользователя
func SetUserBlockHandler(w http.ResponseWriter, r *http.Request) {
	token, _ := getToken(r)
	targetID := r.URL.Query().Get("target_id")

	var body struct {
		IsBlocked bool `json:"is_blocked"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := main_module.SetUserBlockStatus(token, targetID, body.IsBlocked)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}