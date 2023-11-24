package redisdb

import (
	"job-application-api/config"

	"github.com/go-redis/redis"
)

func InitRedisClient(cfg config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDb,
	})
	return client
}
