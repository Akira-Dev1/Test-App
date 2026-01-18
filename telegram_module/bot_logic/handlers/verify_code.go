package handlers

import (
	"bot_logic/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"bot_logic/auth"
)

type VerifyCodeRequest struct {
	ChatId string    `json:"chat_id"`
	Code   string `json:"code"`
}

func VerifyCode(w http.ResponseWriter, r *http.Request) {
	var req VerifyCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(req.ChatId)
	if err != nil {
		http.Error(w, "invalid code", http.StatusBadRequest)
		return
	}
	status := storage.GetUserStatus(id)

	if status == "authorized" {
		err := auth.VerifyCode(id, req.Code)
		if err != nil {
			w.Write([]byte("Error verify code (HANDLER)"))
		} else {
			
			w.Write([]byte("You were successfully logged in. You can continue on your new device!"))
		}
	} else {
		w.Write([]byte("You are not allowed to do this command. Please, sign up/in at first."))
	}
}
