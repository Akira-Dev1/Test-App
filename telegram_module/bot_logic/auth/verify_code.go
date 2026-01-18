package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"bot_logic/storage"
)

type VerifyCodeResponse struct {
	Status string `json:"status"`
}

func VerifyCode(chat_id int, code string) error {
	refresh_token, err := storage.GetRefreshToken(chat_id)
	if err != nil {
		log.Println("Couldn-t find refresh token: ", err.Error())
		return err
	}
	log.Println(refresh_token)

	url := "http://auth:8081/login/verify"

	body := map[string]string{"code": code, "refresh_token": refresh_token}
	jsonData, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Err request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Err response: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return err
	}

	var verifyCodeResp VerifyCodeResponse
	if err := json.Unmarshal(bodyBytes, &verifyCodeResp); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return err
	}

	if verifyCodeResp.Status == "denied" {
		return http.ErrBodyNotAllowed
	}
	if verifyCodeResp.Status == "pending"{
		log.Println("status pending....")
	}
	if verifyCodeResp.Status == "approved"{
		log.Println("status approved....")
	}
	return nil
}
