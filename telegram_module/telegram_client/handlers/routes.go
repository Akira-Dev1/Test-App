package handlers

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func HandleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, update *tgbotapi.Update) {
	if msg.Text == "/start" {
		StartBot(bot, msg)
	} else {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Нет такой команды"))
		return
	}
}

func HandleCallback(bot *tgbotapi.BotAPI, msg *tgbotapi.CallbackQuery, update *tgbotapi.Update) {
	callback := update.CallbackQuery
	data := callback.Data

	answer := tgbotapi.NewCallback(callback.ID, "")
	bot.AnswerCallbackQuery(answer)

	if data == "login" {
		LoginBot(bot, msg)
	} else if data == "github_login" {
		GithubLoginBot(bot, msg)
	} else if data == "yandex_login" {
		YandexLoginBot(bot, msg)
	} else if data == "code_au" {
		CodeLoginBot(bot, msg)
	}

}
