package main

import (	
	"telegram_client/bot"

	"telegram_client/update"
)

func main() {
	go update.UpdateUserStatus()
	go update.GetUserNotifications()

	bot.Run()
}