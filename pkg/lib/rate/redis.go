package rate

import (
	"github.com/go-redis/redis"
)

type RedisRateLimiter struct {
	dao *RedisDAO
}

func NewRedisRateLimiter(config *redis.Options) *RedisRateLimiter {
	return &RedisRateLimiter{
		dao: NewRedisDAO(config),
	}
}

func (l *RedisRateLimiter) IsAllow(ip string) (bool, error) {
	return true, nil
}
