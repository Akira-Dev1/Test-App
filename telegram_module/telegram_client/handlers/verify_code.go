package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func VerifyCode(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, status string) {
	if status == "authorized" {
		url := "http://tg_nginx/login/verify"

		data := map[string]string{"chat_id": fmt.Sprintf("%d", msg.Chat.ID), "code": msg.Text}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Ошибка создания JSON: %v\n", err)
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Ошибка создания запроса: %v\n", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Ошибка запроса: %v\n", err)
			return
		}

		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		m := tgbotapi.NewMessage(msg.Chat.ID, string(body))
		_, err = bot.Send(m)
		if err != nil {
			log.Printf("Error code n: %v", err)
		}

		delete := tgbotapi.NewDeleteMessage(msg.Chat.ID, msg.MessageID)
		bot.Send(delete)

		
	} else {
		m := tgbotapi.NewMessage(msg.Chat.ID, "You are NOT ALLOWED")
		_, err := bot.Send(m)
		if err != nil {
			log.Printf("Error sending messge: %v", err)
		}
	}

}
