package handlers

import (
	// "bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type UserInfoResponse struct {
	UserID   string `json:"user_id"`
	FullName string `json:"full_name"`
}

func GetUserInfoClientHandler(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, status string) {
	if status == "authorized" {
		url := "http://tg_nginx/users/info?chat_id=" + strconv.FormatInt(msg.Chat.ID, 10)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("Ошибка создания запроса: %v\n", err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Ошибка запроса: %v\n", err)
			return
		}

		defer resp.Body.Close()

		var response UserInfoResponse

		body, _ := io.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &response); err != nil {
			log.Printf("Ошибка парсинга JSON: %v\n", err)
			return
		}
		var mess string = "Username: " + response.FullName + " , user id: " + response.UserID

		m := tgbotapi.NewMessage(msg.Chat.ID, mess)
		_, err = bot.Send(m)
		if err != nil {
			log.Printf("Error github n: %v", err)
		}
	} else {
		m := tgbotapi.NewMessage(msg.Chat.ID, "You are already authorized!")
		_, err := bot.Send(m)
		if err != nil {
			log.Printf("Error sending messge: %v", err)
		}
	}
}
