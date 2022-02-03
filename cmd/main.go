package main

import (
	_ "Rest-Api/docs/cmd"
	"Rest-Api/models"
	"Rest-Api/pkg/handler"
	"Rest-Api/pkg/repository"
	"Rest-Api/pkg/services"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Rest-Api
// @version 1.0
// @description This is a Rest-Api server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	db := repository.Init()
	db.AutoMigrate(&models.Post{}, &models.Comment{}, &models.User{})
	repo := repository.NewRepository(*db)

	service := services.NewService(repo)
	e := echo.New()

	routes := handler.NewHandler(service, e).InitRoutes()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(routes.Start("localhost:8080"))

}
