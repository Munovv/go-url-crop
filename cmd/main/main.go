package main

import (
	"context"
	"github.com/Munovv/go-url-crop/pkg/handler"
	"github.com/Munovv/go-url-crop/pkg/repository"
	"github.com/Munovv/go-url-crop/pkg/server"
	"github.com/Munovv/go-url-crop/pkg/service"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db, err := repository.NewPostgresDb(repository.Config{
		Host:     "localhost",
		Port:     "8000",
		Username: "root",
		Password: "",
		Db:       "job_database",
		SslMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("failed initialize database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run("5000", handlers.InitRoutes()); err != nil {
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
