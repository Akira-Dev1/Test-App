package main

import (	
	"telegram_client/bot"

	"telegram_client/update"
)

func main() {
	go update.UpdateUserStatus()

	bot.Run()
}