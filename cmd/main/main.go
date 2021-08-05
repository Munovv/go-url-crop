package main

import (
	"context"
	"github.com/Munovv/go-url-crop/pkg/handler"
	"github.com/Munovv/go-url-crop/pkg/repository"
	"github.com/Munovv/go-url-crop/pkg/server"
	"github.com/Munovv/go-url-crop/pkg/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config: %s", err.Error())
	}

	db, err := repository.NewMysqlDb(repository.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Db:       viper.GetString("database.database"),
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
		if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
