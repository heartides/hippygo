package conf

import (
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

var Cache *redis.Client

func ConfRedis() *redis.Client {
	db, _ := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 64)

	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password:   os.Getenv("REDIS_AUTH"),
		DB:         int(db),
		MaxRetries: 1,
	})
	Cache = client

	return client
}
