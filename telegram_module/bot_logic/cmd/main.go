package main

import (
	"fmt"
	"net/http"

	"bot_logic/handlers"
	"bot_logic/storage"
)

func main(){
	storage.InitRedis()
	handlers.RegisterRoutes()

	fmt.Println("Сервер запущен на :8083")
	http.ListenAndServe(":8083", nil)
}