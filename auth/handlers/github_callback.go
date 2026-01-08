package handlers

import (
	"log"
	"net/http"

	"auth/services"
	"auth/storage"
	"auth/domain"
)

func GithubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	code  := r.URL.Query().Get("code")
	errQ  := r.URL.Query().Get("error")

	if state == "" {
		http.Error(w, "state not found", http.StatusBadRequest)
		return
	}

	if errQ != "" {
		authState, ok := storage.GetAuthState(state)
		if ok {
			authState.Status = domain.StatusDenied
			storage.SaveAuthState(authState, state)
		}
		w.Write([]byte("GitHub authorization denied"))
		return
	}


	// if code == "" {
	// 	http.Error(w, "code not found", http.StatusBadRequest)
	// 	return
	// }

	log.Println("GitHub code:", code) // НЕ ЗАБЫТЬ УБРАТЬ ЛОГ
	

	err := services.GithubAuth(code)
	if err != nil {
		log.Println("GithubAuth error:", err)
		http.Error(w, "github auth failed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("GitHub auth success"))
}

