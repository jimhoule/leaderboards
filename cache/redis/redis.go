package redis

import (
	"main/cache"

	"github.com/redis/go-redis/v9"
)

type Cache = redis.Client
type RedisCacher[TCache Cache] struct{}

func NewRedisCacher() *RedisCacher[Cache] {
	return &RedisCacher[Cache]{}
}

func (rc *RedisCacher[TCache]) New(address string, password string) *cache.CacheHandler[Cache] {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		// NOTE: Use default DB
		DB:       0,
	})


	return &cache.CacheHandler[Cache]{
		Cache: redisClient,
	}
}