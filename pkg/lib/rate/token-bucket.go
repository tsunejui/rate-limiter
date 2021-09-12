package rate

import (
	"sync"

	"golang.org/x/time/rate"
)

type TokenBucketRateLimiter struct {
	options *LimiterOptions
	mu      *sync.RWMutex

	ips map[string]*rate.Limiter
}

func NewTokenBucketRateLimiter(options *LimiterOptions) *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		mu:      &sync.RWMutex{},
		options: options,
		ips:     make(map[string]*rate.Limiter),
	}
}

func (l *TokenBucketRateLimiter) Init() error {
	return nil
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

	limit := rate.Every(l.options.timeRate)
	limiter := rate.NewLimiter(limit, l.options.max)
	l.ips[address] = limiter
	return limiter
}
