package rate

import (
	"github.com/go-redis/redis"
)

type RedisDAO struct {
	client *redis.Client
}

func NewRedisDAO(config *redis.Options) *RedisDAO {
	return &RedisDAO{
		client: redis.NewClient(config),
	}
}
