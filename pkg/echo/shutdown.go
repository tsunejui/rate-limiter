package echo

import (
	"context"
	"os"
	"os/signal"
	"time"

	labstackEcho "github.com/labstack/echo/v4"
)

func Shutdown(server *labstackEcho.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
