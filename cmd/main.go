package main

import (
	"log"

	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"

	"todo"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	svr := new(todo.Server)
	if err := svr.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
