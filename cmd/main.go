package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/Romulq/todo"
	"github.com/Romulq/todo/pdg/handler"
	"github.com/Romulq/todo/pdg/repository"
	"github.com/Romulq/todo/pdg/service"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
