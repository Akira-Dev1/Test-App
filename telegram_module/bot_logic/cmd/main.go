package main

import (
	"fmt"
	"net/http"

	"bot_logic/handlers"
	"bot_logic/storage"
	"bot_logic/update"
)

func main(){
	storage.InitRedis()
	handlers.RegisterRoutes()

	update.RefreshUsersTokens()

	fmt.Println("Сервер запущен на :8083")
	http.ListenAndServe(":8083", nil)
}