package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"fmt"
	"net/http"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func CodeLoginBot(bot *tgbotapi.BotAPI, msg *tgbotapi.CallbackQuery, status string) {
	if status != "authorized" {
		url := "http://tg_nginx/login?type=code"

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
			Code string `json:"code"`
		}
		
		body, _ := io.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &response); err != nil {
			log.Printf("Ошибка парсинга JSON: %v\n", err)
			return
		}

		// КРАСИВОЕ СООБЩЕНИЕ С ЭМОЗДИ ПОМОГЛА ОФОРМИТЬ НЕЙРОСЕТЬ
		messageText := fmt.Sprintf(
            "🔐 *Код для входа*\n\n" +
            "```\n%s\n```\n\n" +
            "📱 *Как использовать:*\n" +
            "1. Откройте приложение с устройства, на котором вы авторизованы\n" +
            "2. Введите команду /verifycode\n" +
			"3. Введите код\n\n" +
            "⏰ *Действует:* 60 секунд\n" +
            "⚠️ *Не сообщайте код никому!*",
            response.Code,
        )

		m := tgbotapi.NewMessage(msg.Message.Chat.ID, messageText)
		_, err = bot.Send(m)
		if err != nil {
			log.Printf("Error code au n: %v", err)
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
