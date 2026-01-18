package storage

import (
    "context"
    "fmt"
    "log"
    "time"
    "github.com/redis/go-redis/v9"
)

var (
    AuthRedis  *redis.Client
)

func InitRedis() error {    
    AuthRedis = redis.NewClient(&redis.Options{
        Addr:     "tg_redis:6380",
        Password: "MaksimKarmyshakovGay",
        DB:       0,
    })
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    if err := AuthRedis.Ping(ctx).Err(); err != nil {
        return fmt.Errorf("Error auth Redis connection: %v", err)
    }
    
    log.Println("Redis connected")
    return nil
}

func CloseRedis() {
    if AuthRedis != nil {
        AuthRedis.Close()
    }
    log.Println("Redis connection closed")
}