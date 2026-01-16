package handlers

import ( 
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
)

func LoginBot(bot *tgbotapi.BotAPI, msg *tgbotapi.CallbackQuery) {
	m := tgbotapi.NewMessage(msg.Message.Chat.ID, "Выберите способ входа: ")

	github_login := tgbotapi.NewInlineKeyboardButtonData("Github OAuth", "github_login")
	yandex_login := tgbotapi.NewInlineKeyboardButtonData("Yandex OAuth", "yandex_login")
	code_au := tgbotapi.NewInlineKeyboardButtonData("Code Auth", "code_au")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(github_login),
		tgbotapi.NewInlineKeyboardRow(yandex_login),
		tgbotapi.NewInlineKeyboardRow(code_au),
	)
	m.ReplyMarkup = keyboard

	_, err := bot.Send(m)
	if err != nil {
		log.Printf("Error sending login button: %v", err)
	}
	delete := tgbotapi.NewDeleteMessage(msg.Message.Chat.ID, msg.Message.MessageID)
	bot.Send(delete)
}