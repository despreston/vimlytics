package redis

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"log"
	"os"
	"sync"
)

var once sync.Once
var Ctx = context.Background()
var c *redis.Client

const Empty = redis.Nil

func Client() *redis.Client {
	once.Do(func() {
		url := os.Getenv("REDIS_URL")

		if len(url) < 1 {
			url = "localhost:6379"
		}

		c = redis.NewClient(&redis.Options{
			Addr:     url,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		_, err := c.Ping(Ctx).Result()

		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
	})

	return c
}
