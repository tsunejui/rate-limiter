package main

import (
	"rate-limiter/middleware"

	"github.com/labstack/echo/v4"
)

func IPRateLimit() echo.MiddlewareFunc {
	return middleware.Limiter(&middleware.MiddlewareCofig{
		Mode: middleware.GeneralMode,
	})
}
