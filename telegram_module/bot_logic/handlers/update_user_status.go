package handlers

import (
	"net/http"
	"log"

	"bot_logic/update"
)

func UpdateUserStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateUserStatusHandler called")
	update.UpdateUserStatus()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update process completed"))
}
