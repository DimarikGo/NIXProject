package main

import (
	"Rest-Api"
	"Rest-Api/models"
	"Rest-Api/pkg/handler"
	"Rest-Api/pkg/repository"
	"Rest-Api/pkg/services"
	"fmt"

	_ "github.com/lib/pq"
	"log"
)

func main() {

	gorm := repository.Init()
	err := gorm.AutoMigrate(&models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("models no migrate")
	}
	repo := repository.NewRepository(*gorm)
	service := services.NewService(repo)
	handlers := handler.NewHandler(service)

	server := new(Rest_Api.Server)
	err = server.Run("8080", handlers.InitRoutes())
	fmt.Println("Rest-Api started")
	if err != nil {
		log.Fatalf("server doesn`t started:%s", err.Error())
	}
}
