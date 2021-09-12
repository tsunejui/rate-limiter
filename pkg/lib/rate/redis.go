package rate

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisRateLimiter struct {
	dao     *RedisDAO
	options *LimiterOptions
}

func NewRedisRateLimiter(config *LimiterOptions) *RedisRateLimiter {
	return &RedisRateLimiter{
		options: config,
	}
}

func (l *RedisRateLimiter) SetRedisOption(config *redis.Options) *RedisRateLimiter {
	l.dao = NewRedisDAO(config)
	return l
}

func (l *RedisRateLimiter) Init() error {
	if err := l.dao.Ping(); err != nil {
		return fmt.Errorf("error init redis: %v", err)
	}
	return nil
}

func (l *RedisRateLimiter) IsAllow(ip string) (bool, error) {
	now := time.Now().UnixNano()
	if err := l.dao.RemoveDataByIPAndScoreRange(ip, 0, now-l.options.timeRate.Nanoseconds()); err != nil {
		return false, fmt.Errorf("failed to remove data by dao: %v", err)
	}

	count, err := l.dao.GetDataCountByIPAndScoreRange(ip, 0, -1)
	if err != nil {
		return false, fmt.Errorf("failed to get data by dao: %v", err)
	}

	isAllow := count < l.options.max
	if isAllow {
		if err := l.dao.CreateData(ip, redis.Z{
			Score:  float64(now),
			Member: float64(now),
		}); err != nil {
			return false, fmt.Errorf("failed to add key by dao: %v", err)
		}

		if err := l.dao.SetExpire(ip, l.options.timeRate); err != nil {
			return false, fmt.Errorf("failed to set the expiration by dao: %v", err)
		}
	}

	return isAllow, nil
}
