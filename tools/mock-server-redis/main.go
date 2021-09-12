package main

import (
	"fmt"
	"net/http"
	"rate-limiter/tools/mock-server-redis/conf"
	"time"

	rMiddleware "rate-limiter/middleware"
	pkgEcho "rate-limiter/pkg/echo"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

var (
	timeRate   = time.Minute * 1
	maxRequest = 5
)

const cacheDB = 1

func main() {
	env := conf.GetConfig()
	e := echo.New()
	e.Use(
		rMiddleware.Limiter(&rMiddleware.MiddlewareCofig{
			Mode:     rMiddleware.RedisMode,
			TimeRate: timeRate,
			Max:      maxRequest,
			RedisOptions: &redis.Options{
				Addr:     env.RedisAddress,
				Password: env.RedisPassword,
				DB:       cacheDB,
			},
		}),
	)
	e.GET("/", hello)

	go func() {
		e.Logger.Fatal(
			e.Start(
				fmt.Sprintf(":%s", env.ListenPort),
			),
		)
	}()

	pkgEcho.Shutdown(e)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
