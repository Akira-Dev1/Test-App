package bot

import (
	"log"
	"os"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type App struct {
	Bot *tgbotapi.BotAPI
}

func Run() {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	app := &App{Bot: bot}

	app.Router()
}
