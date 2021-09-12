package rate

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type TokenBucketRateLimiter struct {
	mu       *sync.RWMutex
	timeRate time.Duration
	max      int
	ips      map[string]*rate.Limiter
}

func NewTokenBucketRateLimiter(timeRate time.Duration, m int) *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		timeRate: timeRate,
		max:      m,
		mu:       &sync.RWMutex{},
		ips:      make(map[string]*rate.Limiter),
	}
}

func (l *TokenBucketRateLimiter) IsAllow(ip string) (bool, error) {
	l.mu.Lock()
	limiter, exists := l.ips[ip]
	l.mu.Unlock()

	if !exists {
		limiter = l.addIP(ip)
	}
	return limiter.Allow(), nil
}

func (l *TokenBucketRateLimiter) addIP(address string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()

	limit := rate.Every(1 * time.Minute)
	limiter := rate.NewLimiter(limit, l.max)
	l.ips[address] = limiter
	return limiter
}
