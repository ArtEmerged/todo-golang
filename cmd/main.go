package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"

	"todo"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewSqliteDB(&repository.Config{
		Driver: viper.GetString("db.driver"),
		Dsn:    viper.GetString("db.dsn"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	svr := new(todo.Server)
	go func() {
		if err := svr.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Println("TodoApp Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("TodoApp Shutting Down")
	if err := svr.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down : %s", err.Error())
	}
	if err:= db.Close(); err != nil {
		log.Printf("error occured on db connection close : %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
