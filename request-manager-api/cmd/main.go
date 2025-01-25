package main

import (
	"log"
	"request_manager_api"
	"request_manager_api/pkg/handlers"
)

func main() {
	handlers := new(handlers.Handlers)
	srv := new(request_manager_api.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
