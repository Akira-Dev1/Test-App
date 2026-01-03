package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	"fmt"
	"auth/domain"
	"auth/storage"
)

type LoginRequest struct {
	EntryToken string `json:"entry_token"`
}

type CodeResponse struct {
	Code string `json:"code"`
}


func Login(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр type
	loginType := r.URL.Query().Get("type")
	
	if loginType != "code" {
		http.Error(w, "unsupported login type", http.StatusBadRequest)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EntryToken == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// 1. сохраняем состояние авторизации
	authState := domain.AuthState{
		EntryToken: req.EntryToken,
		Status:     domain.StatusPending,
		ExpiresAt:  time.Now().Add(5 * time.Minute),
	}
	storage.SaveAuthState(authState)

	// 2. генерируем код
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000

	codeState := domain.CodeState{
		Code: fmt.Sprintf("%06d", code),
		EntryToken: req.EntryToken,
		ExpiresAt:  time.Now().Add(1 * time.Minute),
	}
	fmt.Println("Saving code:", codeState)

	storage.SaveCode(codeState)

	// 3. отдаём код клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CodeResponse{
		Code: codeState.Code,
	})
}

func VerifyCode(w http.ResponseWriter, r *http.Request) {
    // 1. Декодируем JSON из тела запроса
    var req domain.VerifyCodeRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EntryToken == "" || req.Code == "" {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    // 2. Получаем код по entry_token
    codeState, ok := storage.GetCode(req.EntryToken)
    if !ok {
        // Обновляем authState в denied, если он есть
        if authState, exists := storage.GetAuthState(req.EntryToken); exists {
            authState.Status = domain.StatusDenied
            storage.SaveAuthState(authState)
        }
        http.Error(w, "code not found", http.StatusNotFound)
        return
    }

    // 3. Проверяем TTL
    if time.Now().After(codeState.ExpiresAt) {
        if authState, exists := storage.GetAuthState(req.EntryToken); exists {
            authState.Status = domain.StatusDenied
            storage.SaveAuthState(authState)
        }
        http.Error(w, "code expired", http.StatusForbidden)
        return
    }

    // 4. Сравниваем коды
    if codeState.Code != req.Code {
        if authState, exists := storage.GetAuthState(req.EntryToken); exists {
            authState.Status = domain.StatusDenied
            storage.SaveAuthState(authState)
        }
        http.Error(w, "invalid code", http.StatusForbidden)
        return
    }

    // 5. Меняем статус авторизации на approved
    authState, ok := storage.GetAuthState(req.EntryToken)
    if !ok {
        http.Error(w, "auth state not found", http.StatusNotFound)
        return
    }
    authState.Status = domain.StatusApproved
    storage.SaveAuthState(authState)

    // 6. Отдаём ответ клиенту
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(domain.VerifyCodeResponse{
        Status: string(authState.Status),
    })
}

// func VerifyCode(w http.ResponseWriter, r *http.Request) {
//     // 1. Декодируем JSON из тела запроса
//     var req domain.VerifyCodeRequest
//     if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EntryToken == "" || req.Code == "" {
//         http.Error(w, "invalid request", http.StatusBadRequest)
//         return
//     }

//     // 2. Получаем код по entry_token

// 	fmt.Println("Looking up code for:", req.EntryToken)
// 	fmt.Println("All codes in storage:", storage.AllCodes())


//     codeState, ok := storage.GetCode(req.EntryToken)
//     if !ok {
//         http.Error(w, "code not found", http.StatusNotFound)
//         return
//     }

//     // 3. Проверяем TTL
//     if time.Now().After(codeState.ExpiresAt) {
//         http.Error(w, "code expired", http.StatusForbidden)
//         return
//     }

//     // 4. Сравниваем коды
//     if codeState.Code != req.Code {
//         http.Error(w, "invalid code", http.StatusForbidden)
//         return
//     }

//     // 5. Меняем статус авторизации
//     authState, ok := storage.GetAuthState(req.EntryToken)
//     if !ok {
//         http.Error(w, "auth state not found", http.StatusNotFound)
//         return
//     }
//     authState.Status = domain.StatusApproved
//     storage.SaveAuthState(authState)

//     // 6. Отдаём ответ клиенту
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(domain.VerifyCodeResponse{
//         Status: string(authState.Status),
//     })
// }
