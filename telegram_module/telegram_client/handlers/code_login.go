package handlers

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
)

func CodeLoginBot(bot *tgbotapi.BotAPI, msg *tgbotapi.CallbackQuery, status string) {
	m := tgbotapi.NewMessage(msg.Message.Chat.ID, "CODE")
	_, err := bot.Send(m)
	if err != nil {
		log.Printf("Error code n: %v", err)
	}
	delete := tgbotapi.NewDeleteMessage(msg.Message.Chat.ID, msg.Message.MessageID)
	bot.Send(delete)
}
