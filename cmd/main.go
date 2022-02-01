package main

import (
	"Rest-Api/pkg/handler"
	"Rest-Api/pkg/repository"
	"Rest-Api/pkg/services"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	db := repository.Init()
	repo := repository.NewRepository(*db)

	service := services.NewService(repo)
	e := echo.New()

	routes := handler.NewHandler(service, e).InitRoutes()

	e.Logger.Fatal(routes.Start("localhost:8080"))

}
