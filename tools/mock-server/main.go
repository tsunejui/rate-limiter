package main

import (
	"fmt"
	"net/http"
	"rate-limiter/tools/mock-server/conf"
	"time"

	rMiddleware "rate-limiter/middleware"
	pkgEcho "rate-limiter/pkg/echo"

	"github.com/labstack/echo/v4"
)

var (
	timeRate   = time.Minute * 5
	maxRequest = 5
)

func main() {
	e := echo.New()
	e.Use(
		rMiddleware.Limiter(&rMiddleware.MiddlewareCofig{
			Mode:     rMiddleware.GeneralMode,
			TimeRate: timeRate,
			Max:      maxRequest,
		}),
	)
	e.GET("/", hello)

	go func() {
		e.Logger.Fatal(
			e.Start(
				fmt.Sprintf(":%s", conf.GetConfig().ListenPort),
			),
		)
	}()

	pkgEcho.Shutdown(e)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
