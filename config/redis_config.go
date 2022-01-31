package config

import (
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func StartRedisConfig()  {
	
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})
}