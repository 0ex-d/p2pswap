package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	log.Infof("Connecting to Redis")
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6888",
		DB:   0,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
	RedisClient.Set(context.Background(), "health_check", time.Now().Unix(), 0)
}
