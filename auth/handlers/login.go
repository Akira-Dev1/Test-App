package handlers

import (
	"encoding/json"
	"net/http"

	"net/url"
	"os"

	"auth/codeauth"
	"time"

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

	if r.URL.Query().Get("type") == "github" {
		clientID := os.Getenv("GITHUB_CLIENT_ID")
		if clientID == "" {
			http.Error(w, "github client id not set", http.StatusInternalServerError)
			return
		}

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EntryToken == "" {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		authState := domain.AuthState{
			Status:     domain.StatusPending,
			ExpiresAt:  time.Now().Add(5 * time.Minute),
		}
		storage.SaveAuthState(authState, req.EntryToken)

		params := url.Values{}
		params.Add("client_id", clientID)
		params.Add("state", req.EntryToken)
		params.Add("scope", "user:email") // строка, описывающая, к каким данным пользователь разрешает доступ приложению

		githubAuthURL := "https://github.com/login/oauth/authorize?" + params.Encode()

		http.Redirect(w, r, githubAuthURL, http.StatusFound)



	} else if r.URL.Query().Get("type") == "yandex" {
		clientID := os.Getenv("YANDEX_CLIENT_ID")
		if clientID == "" {
			http.Error(w, "yandex client id not set", http.StatusInternalServerError)
			return
		}

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EntryToken == "" {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		authState := domain.AuthState{
			Status:     domain.StatusPending,
			ExpiresAt:  time.Now().Add(5 * time.Minute),
		}
		storage.SaveAuthState(authState, req.EntryToken)

		params := url.Values{}
		params.Add("client_id", clientID)
		params.Add("response_type", "code")  // ОБЯЗАТЕЛЬНО для Яндекс
		params.Add("state", req.EntryToken)
		params.Add("redirect_uri", os.Getenv("YANDEX_REDIRECT_URI"))  // Обязательно
		params.Add("scope", "login:email login:info")  // Яндекс scopes, а не GitHub

		yandexAuthURL := "https://oauth.yandex.ru/authorize?" + params.Encode()

		http.Redirect(w, r, yandexAuthURL, http.StatusFound)



	} else if r.URL.Query().Get("type") == "code" {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EntryToken == "" {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		// генерируем код
		code, err := codeauth.GenerateCode(req.EntryToken) // здесь респонс от компонента Code Autentification - он возвр. код

		// создаём auth state
		authState := domain.AuthState{
			Status:     domain.StatusPending,
			ExpiresAt:  time.Now().Add(5 * time.Minute),
		}
		storage.SaveAuthState(authState, req.EntryToken)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(CodeResponse{Code: code})


	} else {
		http.Error(w, "unsupported login type", http.StatusBadRequest)
		return
	}
}



// else if r.URL.Query().Get("type") == "yandex" {
//     clientID := os.Getenv("YANDEX_CLIENT_ID")
//     if clientID == "" {
//         http.Error(w, "yandex client id not set", http.StatusInternalServerError)
//         return
//     }

//     var req LoginRequest
//     if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EntryToken == "" {
//         http.Error(w, "invalid request", http.StatusBadRequest)
//         return
//     }

//     authState := domain.AuthState{
//         Status:     domain.StatusPending,
//         ExpiresAt:  time.Now().Add(5 * time.Minute),
//     }
//     storage.SaveAuthState(authState, req.EntryToken)


// }