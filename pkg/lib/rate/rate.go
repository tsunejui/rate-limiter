package rate

import "time"

type Limiter interface {
	Init() error
	IsAllow(ip string) (bool, error)
}

type LimiterOptions struct {
	max      int
	timeRate time.Duration
}

func NewLimiterOptions(timeRate time.Duration, max int) *LimiterOptions {
	return &LimiterOptions{
		timeRate: timeRate,
		max:      max,
	}
}
