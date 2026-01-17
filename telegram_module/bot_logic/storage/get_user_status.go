package storage

import (
	"context"
	"strconv"
)

var ctx = context.Background()

func GetUserStatus(chat_id int)(string){
	id := strconv.Itoa(chat_id)
	session, err := AuthRedis.Exists(ctx, id).Result()
	if err != nil {
		panic(err)
	}
	if session == 0 {
		return "unknown"
	} else {
		status, err := AuthRedis.HGet(ctx, id, "status").Result()
		if err != nil {
			panic(err)
		}

		return status
	}
}
