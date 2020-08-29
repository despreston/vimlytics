package cache

import (
	"github.com/despreston/vimlytics/redis"
	"log"
	"time"
)

const ttl = 72 * time.Hour

func Get(key string) (string, bool) {
	var val string

	val, err := redis.Client().Get(redis.Ctx, key).Result()

	if err == redis.Empty {
		return "", false
	} else if err != nil {
		log.Printf("Error @ redis GET: %v", err)
		return "", false
	}

	return val, true
}

func Set(key string, val string) {
	_, err := redis.Client().SetNX(redis.Ctx, key, val, ttl).Result()

	if err != nil {
		log.Printf("Error @ redis SET: %v", err)
	}
}
