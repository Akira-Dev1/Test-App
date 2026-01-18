package storage

import(
	"github.com/redis/go-redis/v9"
)

func GetAuthRedis() *redis.Client {
    return AuthRedis
}