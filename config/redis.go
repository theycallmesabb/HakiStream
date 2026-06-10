package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func ConnectRedis() {
	var ctx = context.Background()
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis connection failed: ", err)
	}

	log.Println("Redis is connected")

}
