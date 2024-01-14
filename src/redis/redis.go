package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func ConnectRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetCache(key string) string {
	ctx := context.Background()

	jsonString, _ := client.Get(ctx, key).Result()
	return jsonString
}

func SetCache(key string, data string) error {
	ctx := context.Background()
	return client.Set(ctx, key, data, 0).Err()
}
