package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var emailCodes = map[string]string{} // временное хранилище кодов на email

type LoginRequest struct {
	Email string `json:"email"`
}

// RequestLoginCode отправляет код подтверждения на email
func RequestLoginCode(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Генерация 6-значного кода
	rand.Seed(time.Now().UnixNano())
	code := GenerateCode()
	emailCodes[req.Email] = code

	// Отправка письма
	err := sendEmail(req.Email, code)
	if err != nil {
		http.Error(w, "failed to send email", http.StatusInternalServerError)
		fmt.Println("Email send error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "code sent"})
}

// VerifyLoginCode проверяет введённый пользователем код
type VerifyRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func VerifyLoginCode(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if code, ok := emailCodes[req.Email]; !ok || code != req.Code {
		http.Error(w, "invalid code", http.StatusUnauthorized)
		return
	}

	// TODO: здесь можно создать или обновить пользователя в MongoDB
	// TODO: сгенерировать токен сессии (JWT) и вернуть его клиенту

	delete(emailCodes, req.Email) // код больше не нужен

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "login successful"})
}
