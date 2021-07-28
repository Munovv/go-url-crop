package main

import (
	"context"
	"github.com/Munovv/go-url-crop/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := new(server.Server)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run("5000"); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Print("Server started")

	<-done
	log.Print("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error occured on server shutdown: %s", err.Error())
	}
	log.Print("Server exited properly")
}
