package config

import (
	"fmt"

	"github.com/go-redis/redis"
)

func StartRedisConfig() *redis.Client {
	
	RedisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

	pong, err := RedisClient.Ping().Result()

	fmt.Println(pong, err)

	return RedisClient
}