package handlers

import (
	"bot_logic/auth"
	"bot_logic/storage"
	"log"

	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	ChatId int `json:"chat_id"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Get("type") == "github" {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		status := storage.GetUserStatus(req.ChatId)

		if status == "anonymous" {
			if err := storage.UpdateUser(req.ChatId); err != nil {
				http.Error(w, "Couldn-t update user", http.StatusBadRequest)
				return
			}
		} else if status == "unknown" {
			err := storage.CreateUser(req.ChatId)
			if err != nil {
				http.Error(w, "Couldn-t create user", http.StatusBadRequest)
				return
			}
		}

		entry_token, err := storage.GetEntryToken(req.ChatId)
		if err != nil {
			http.Error(w, "Couldn-t get entry token", http.StatusBadRequest)
			return
		}

		url, err := auth.GithubLogin(req.ChatId, entry_token)
		if err != nil {
			http.Error(w, "Couldn-t get entry token 2", http.StatusBadRequest)
			return
		}
		log.Println(url, " url")

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"url": url,
		})

	} else if r.URL.Query().Get("type") == "yandex" {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		status := storage.GetUserStatus(req.ChatId)

		if status == "anonymous" {
			if err := storage.UpdateUser(req.ChatId); err != nil {
				http.Error(w, "Couldn-t update user", http.StatusBadRequest)
				return
			}
		} else if status == "unknown" {
			err := storage.CreateUser(req.ChatId)
			if err != nil {
				http.Error(w, "Couldn-t create user", http.StatusBadRequest)
				return
			}
		}

		entry_token, err := storage.GetEntryToken(req.ChatId)
		if err != nil {
			http.Error(w, "Couldn-t get entry token", http.StatusBadRequest)
			return
		}

		url, err := auth.YandexLogin(req.ChatId, entry_token)
		if err != nil {
			http.Error(w, "Couldn-t get entry token 2", http.StatusBadRequest)
			return
		}
		log.Println(url, " url")

		w.Header().Set("Content-Type", "application/json")
		
		json.NewEncoder(w).Encode(map[string]string{
			"url": url,
		})
	} else if r.URL.Query().Get("type") == "code" {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		auth.CodeAuth(req.ChatId)
	}
}
