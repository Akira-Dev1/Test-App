package handlers

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
)

func GithubLoginBot(bot *tgbotapi.BotAPI, msg *tgbotapi.CallbackQuery) {
	
	m := tgbotapi.NewMessage(msg.Message.Chat.ID, "Github")
	_, err := bot.Send(m)
	if err != nil {
		log.Printf("Error github n: %v", err)
	}
	delete := tgbotapi.NewDeleteMessage(msg.Message.Chat.ID, msg.Message.MessageID)
	bot.Send(delete)
}
