package handlers

import (
	"net/http"
	"bot_logic/storage"

	"encoding/json"
)

type GetUserStatusRequest struct {
	ChatId int `json:"chat_id"`
}

type GetUserStatusResponse struct {
	Status string `json:"status"`
}


func GetUserStatus(w http.ResponseWriter, r *http.Request) () {
	var req GetUserStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	status := storage.GetUserStatus(req.ChatId)

	json.NewEncoder(w).Encode(GetUserStatusResponse{Status: status})

}