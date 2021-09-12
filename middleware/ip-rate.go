package middleware

import (
	"fmt"
	"net/http"
	"rate-limiter/pkg/lib/rate"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	limiter rate.Limiter
}

const (
	RedisMode   = "redis-mode"
	GeneralMode = "general-mode"
)

type MiddlewareCofig struct {
	Mode         string
	TimeRate     time.Duration
	Max          int
	RedisOptions *redis.Options
}

func Limiter(config *MiddlewareCofig) echo.MiddlewareFunc {
	var limiter rate.Limiter
	switch config.Mode {
	case GeneralMode:
		limiter = rate.NewTokenBucketRateLimiter(config.TimeRate, config.Max)
	default:
		panic(
			fmt.Sprintf("invalid mode: %s", config.Mode),
		)
	}
	m := &Middleware{
		limiter: limiter,
	}
	return m.MiddlewareFunc
}

func (m *Middleware) MiddlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()
		allow, err := m.limiter.IsAllow(ip)
		if err != nil {
			c.Error(err)
			return fmt.Errorf("failed to use limiter: %v", err)
		}

		if !allow {
			return echo.NewHTTPError(http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
		}

		if err := next(c); err != nil {
			c.Error(err)
			return fmt.Errorf("failed to continue: %v", err)
		}
		return nil
	}
}
