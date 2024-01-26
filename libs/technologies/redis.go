package technologies

import (
	"context"
	"strconv"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Host     string `envconfig:"HOST"`
	Port     int    `envconfig:"PORT"`
	DB       int    `envconfig:"DB"`
	Password string `envconfig:"PASSWORD"`
}

func InitRedis(rCfg RedisConfig) *redis.Client {
	address := rCfg.Host + ":" + strconv.Itoa(rCfg.Port)
	clientRedis := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: rCfg.Password, // no password set
		DB:       rCfg.DB,
	})

	ctx := context.Background()
	clientRedis.Do(ctx, "PING")

	return clientRedis
}
