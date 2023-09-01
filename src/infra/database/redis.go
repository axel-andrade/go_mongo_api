package database

import (
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func ConnectRedisDB() {
	//Initializing redis
	dsn := os.Getenv("REDIS_HOST")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	password := os.Getenv("REDIS_PASSWORD")

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		db = 0
	}

	client = redis.NewClient(&redis.Options{
		Addr:     dsn,      //redis port
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	var ctx = context.TODO()

	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisDB() *redis.Client {
	return client
}
