package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"rate-limiter/tools/mock-server/conf"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)

	go func() {
		e.Logger.Fatal(
			e.Start(
				fmt.Sprintf(":%s", conf.GetConfig().ListenPort),
			),
		)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
