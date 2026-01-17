package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func YandexLoginBot(bot *tgbotapi.BotAPI, msg *tgbotapi.CallbackQuery, status string) {
	if status != "authorized" {
		url := "http://tg_nginx/login?type=yandex"

		data := map[string]int64{"chat_id": msg.Message.Chat.ID}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Ошибка создания JSON: %v\n", err)
			return
		}

		req, _ := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Ошибка запроса: %v\n", err)
			return
		}

		defer resp.Body.Close()

		var response struct {
			URL   string `json:"url"`
		}
		
		body, _ := io.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &response); err != nil {
			log.Printf("Ошибка парсинга JSON: %v\n", err)
			return
		}

		m := tgbotapi.NewMessage(msg.Message.Chat.ID, response.URL)
		_, err = bot.Send(m)
		if err != nil {
			log.Println(response.URL)
			log.Printf("Error yandex n: %v", err)
		}
	} else {
		m := tgbotapi.NewMessage(msg.Message.Chat.ID, "You are already authorized!")
		_, err := bot.Send(m)
		if err != nil {
			log.Printf("Error sending messge: %v", err)
		}
	}
	delete := tgbotapi.NewDeleteMessage(msg.Message.Chat.ID, msg.Message.MessageID)
	bot.Send(delete)
}
