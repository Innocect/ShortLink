package redisDao

import (
	"github.com/go-redis/redis"
)

type ShortenUrl struct {
	ShortUrl string `json:"shorturl"`
	LongUrl  string `json:"longurl"`
}

func RedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}
