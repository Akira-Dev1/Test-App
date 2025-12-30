package handlers

import (
	"encoding/json"
	"net/http"
	"log"
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

func LoginEmail(w http.ResponseWriter, r *http.Request) {
    type request struct {
        Email string `json:"email"`
    }

    var req request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

    if req.Email == "" {
        http.Error(w, "email required", http.StatusBadRequest)
        return
    }

    code := "123456" // stub, потом random
    log.Println("Send code", code, "to email", req.Email)

    // TODO: save code to Redis with TTL

    json.NewEncoder(w).Encode(map[string]string{
        "message": "code sent to email",
    })
}
