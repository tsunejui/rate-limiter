package rate

import (
	"fmt"
	"time"

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

func (r *RedisDAO) Ping() error {
	_, err := r.client.Ping().Result()
	if err != nil {
		return fmt.Errorf("failed to ping the redis db: %v", err)
	}
	return nil
}

func (r *RedisDAO) SetExpire(key string, expireTime time.Duration) error {
	if err := r.client.Expire(key, expireTime).Err(); err != nil {
		return fmt.Errorf("failed to set the expiration: %v", err)
	}
	return nil
}

func (r *RedisDAO) CreateData(key string, z redis.Z) error {
	if err := r.client.ZAddNX(key, z).Err(); err != nil {
		return fmt.Errorf("failed to create key: %v", err)
	}
	return nil
}

func (r *RedisDAO) GetDataCountByIPAndScoreRange(ip string, min, max int64) (int, error) {
	req, err := r.client.ZRange(ip, min, max).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get values: %v", err)
	}
	return len(req), nil
}

func (r *RedisDAO) RemoveDataByIPAndScoreRange(ip string, min, max int64) error {
	minScore := fmt.Sprint(min)
	maxScore := fmt.Sprint(max)
	if err := r.client.ZRemRangeByScore(ip, minScore, maxScore).Err(); err != nil {
		return fmt.Errorf("failed to remove vaules: %v", err)
	}
	return nil
}
