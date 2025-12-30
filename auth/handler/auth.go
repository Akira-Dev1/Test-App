package handler

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"entry_token": "stub-entry-token",
	}
	json.NewEncoder(w).Encode(resp)
}

func Check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("check stub"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout stub"))
}
