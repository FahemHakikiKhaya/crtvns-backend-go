package main

import (
	"fmt"
	"net/http"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/config"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/controller"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/helper"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/repository"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/router"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/service"
)

func main() {
	fmt.Println("Start server")

	// database
	db := config.DatabaseConnection()

	// repository
	userRepository := repository.NewUserRepository(db)
	userHistoryRepository := repository.NewUserHistoryRepository(db)

	// service
	authService := service.NewAuthServiceImpl(userRepository)
	userHistoryService := service.NewUserHistoryService(userHistoryRepository)

	// controller
	authController := controller.NewAuthController(authService)
	userHistoryController := controller.NewUserHistoryController(userHistoryService)

	routes := router.NewRouter(authController, userHistoryController)

	server := http.Server{Addr: "localhost: 8888", Handler: routes}

	err := server.ListenAndServe()
	
	helper.PanicIfError(err)
}