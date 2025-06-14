package main

import (
	"context"
	"os"
	"os/signal"
	"request_manager_api"
	"request_manager_api/pkg/handlers"
	"request_manager_api/pkg/repository"
	"request_manager_api/pkg/services"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	db, err := repository.NewMysqlDb(repository.Config{
		Host:     os.Getenv("DB_HOST"),     // "req-db"
		Port:     os.Getenv("DB_PORT"),     // "3306"
		Username: os.Getenv("DB_USER"),     // "root"
		Password: os.Getenv("DB_PASSWORD"), // from Secrets
		Dbname:   os.Getenv("DB_NAME"),     // from Secrets
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := services.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := new(request_manager_api.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error running server: %s", err.Error())
		}
	}()
	logrus.Printf("E-RequestControl started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("E-RequestControl shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
