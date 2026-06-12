package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func ConnectRedis() {
	var ctx = context.Background()
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD")})
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis connection failed: ", err)
	}

	log.Println("Redis is connected")

}
