package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"sync"
)

var once sync.Once
var Ctx = context.Background()
var db *mongo.Database

func Db() *mongo.Database {
	once.Do(func() {
		url := os.Getenv("MONGO_URL")

		if len(url) < 1 {
			url = "mongodb://localhost:27017"
		}

		client, err := mongo.Connect(Ctx, options.Client().ApplyURI(url))
		if err != nil {
			log.Fatalf("Failed to connect to Mongo: %+v", err)
		}

		if err := client.Ping(Ctx, readpref.Primary()); err != nil {
			log.Fatalf("Failed to ping to Redis: %v", err)
		}

		db = client.Database("vimlytics")
	})

	return db
}
