package handlers

import (
	"context"
	"time"
	"log"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

// func InitRedis(addr string) {
// 	rdb = redis.NewClient(&redis.Options{
// 		Addr: addr, 
// 	})
// }


func InitRedis(addr string) {
	rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("REDIS NOT CONNECTED:", err)
	}
	log.Println("REDIS CONNECTED OK")
}


// SetCode сохраняет код в Redis с TTL
func SetCode(email, code string, ttl time.Duration) error {
	log.Printf("REDIS SET: key=%s value=%s ttl=%v\n", email, code, ttl)
	return rdb.Set(ctx, email, code, ttl).Err()
}

// GetCode получает код из Redis
func GetCode(email string) (string, error) {
	return rdb.Get(ctx, email).Result()
}

// DeleteCode удаляет код из Redis
func DeleteCode(email string) error {
	return rdb.Del(ctx, email).Err()
}
