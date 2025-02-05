package main

import (
	"log"
	"request_manager_api"
	"request_manager_api/pkg/handlers"
	"request_manager_api/pkg/repository"
	"request_manager_api/pkg/services"
)

func main() {

	repos := repository.NewRepository()
	services := services.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := new(request_manager_api.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
