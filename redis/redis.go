package redis

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"log"
	"sync"
)

var once sync.Once
var Ctx = context.Background()
var c *redis.Client

const Empty = redis.Nil

func Client() *redis.Client {
	once.Do(func() {
		c = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		_, err := c.Ping(Ctx).Result()

		if err != nil {
			log.Fatal("Failed to connect to Redis: %v", err)
		}
	})

	return c
}
