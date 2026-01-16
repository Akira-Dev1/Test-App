package main

import (
	"net/http"
	"log"
)

func main(){
	log.Println("Auth service started on: 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}