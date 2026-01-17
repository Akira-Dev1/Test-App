package handlers

import (
	"bytes"
	"encoding/json"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
)

type GetUserStatusResponse struct {
	Status string `json:"status"`
}

func HandleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, update *tgbotapi.Update) {

	if msg.Text == "/start" {
		StartBot(bot, msg)
	} else {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Нет такой команды"))
		return
	}
}

func HandleCallback(bot *tgbotapi.BotAPI, msg *tgbotapi.CallbackQuery, update *tgbotapi.Update) {

	url := "http://tg_nginx/get_user_status"

	body := map[string]int64{"chat_id": msg.Message.Chat.ID}
	jsonData, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Ошибка создания запроса: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	stat, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка запроса: %v\n", err)
		return
	}
	defer stat.Body.Close() 

    if stat.StatusCode != http.StatusOK {
        log.Printf("Сервер вернул ошибку! Статус: %s", stat.Status)
        return
    }

	var status GetUserStatusResponse
	if err := json.NewDecoder(stat.Body).Decode(&status); err != nil {
		log.Printf("Ошибка конвертироваия: %v\n", err)
		return
	}

	callback := update.CallbackQuery
	data := callback.Data

	answer := tgbotapi.NewCallback(callback.ID, "")
	bot.AnswerCallbackQuery(answer)

	if data == "login" {
		LoginBot(bot, msg, status.Status)
	} else if data == "github_login" {
		GithubLoginBot(bot, msg, status.Status)
	} else if data == "yandex_login" {
		YandexLoginBot(bot, msg, status.Status)
	} else if data == "code_au" {
		CodeLoginBot(bot, msg, status.Status)
	}
}
