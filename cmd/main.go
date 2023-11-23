package main

import (
	"log"

	"todo"
	"todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	svr := new(todo.Server)
	if err := svr.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
