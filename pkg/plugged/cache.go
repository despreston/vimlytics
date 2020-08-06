package plugged

import (
	"github.com/despreston/vimlytics/redis"
	"log"
	"time"
)

const ttl = 24 * time.Hour

func cacheGet(name string) (Plugin, bool) {
	var description string
	var p Plugin

	description, err := redis.Client().Get(redis.Ctx, name).Result()

	if err == redis.Empty {
		return p, false
	} else if err != nil {
		log.Println("Error @ redis GET: %v", err)
		return p, false
	}

	p.Name = name
	p.Description = description

	return p, true
}

func cacheSet(p Plugin) {
	_, err := redis.Client().SetNX(redis.Ctx, p.Name, p.Description, ttl).Result()

	if err != nil {
		log.Println("Error @ redis SET: %v", err)
	}
}
