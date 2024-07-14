package cache

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"

	"github/ariifysp/go-101/config"
)

type (
	RedisInterface interface {
		Connect() *redis.Client
	}

	RedisCache struct {
		*redis.Client
	}
)

var (
	once          sync.Once
	redisInstance *RedisCache
)

func NewRedisClient(config *config.Redis) RedisInterface {
	once.Do(func() {
		connectionURL := GenerateRedisConnectionURL(config)

		client := redis.NewClient(&redis.Options{
			Addr:     connectionURL,
			Password: "",
			DB:       0,
		})

		err := client.Ping(context.Background()).Err()
		if err != nil {
			panic(err)
		}

		log.Println("Redis has been connected")

		redisInstance = &RedisCache{client}
	})

	return redisInstance
}

func GenerateRedisConnectionURL(config *config.Redis) string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}

func (cache *RedisCache) Connect() *redis.Client {
	return cache.Client
}
