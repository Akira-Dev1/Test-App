package handlers

import ( 
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
)

func StartBot(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	reply := "Доброго времени суток! Вас приветсвует команда KVARTIRKA 31! Для того, чтобы продолжить, авторизуйтесь: "

	loginBtn := tgbotapi.NewInlineKeyboardButtonData("🔑 Войти / Авторизоваться", "login")
		
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(loginBtn),
	)

	mesg := tgbotapi.NewMessage(msg.Chat.ID, reply)
	mesg.ReplyMarkup = keyboard

	if _, err := bot.Send(mesg); err != nil {
		log.Printf("Ошибка: %v", err)
	}
}
