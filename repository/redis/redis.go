package redis

import "github.com/go-redis/redis"

type Redis struct {
	redis *redis.Client
}

// New ...
func New(redis *redis.Client) *Redis {
	return &Redis{redis}
}
