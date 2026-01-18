package handlers

import (
	"net/http"
	"encoding/json"

	"bot_logic/update"
)

func GetNotifications(w http.ResponseWriter, r *http.Request) {
	data, err := update.CollectNotificationsFromAllUsers()
	if err != nil {
        http.Error(w, "Failed to collect notifications", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
}
