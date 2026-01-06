package handlers

import (
	"net/http"
	"net/url"
	"os"
)

func GithubLogin(w http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		http.Error(w, "github client id not set", http.StatusInternalServerError)
		return
	}

	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("scope", "user:email")

	githubAuthURL := "https://github.com/login/oauth/authorize?" + params.Encode()

	http.Redirect(w, r, githubAuthURL, http.StatusFound)
}
