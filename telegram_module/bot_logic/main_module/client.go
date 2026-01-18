package main_module

import (
	"net/http"
	"time"
)

const baseURL = "http://main_module:18080"

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func setAuthHeader(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
}
