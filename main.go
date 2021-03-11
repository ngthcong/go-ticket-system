package main

import (
	"context"
	"go-ticket-system/internal/app"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait = 15 * time.Second
	application := app.New()

	go func() {
		if err := application.Start(); err != nil {
			application.Logger().Info(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	application.Stop(ctx)
	// Optionally, you could run srv.Stop in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	application.Logger().Info("shutting down the server")
	os.Exit(0)
}
