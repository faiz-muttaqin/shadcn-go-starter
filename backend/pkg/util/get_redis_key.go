package util

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func GetRedis(key string, redisDB *redis.Client) string {
	val, _ := redisDB.Get(context.Background(), key).Result()
	return val
}
