package db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func SaveToRedis(shortURL string, originalURL string) error {
	return redisClient.Set(ctx, shortURL, originalURL, 0).Err()
}

func GetFromRedis(shortURL string) (string, error) {
	return redisClient.Get(ctx, shortURL).Result()
}
