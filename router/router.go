package router

import (
	"net/http"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/controller"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func NewRouter(authController *controller.AuthController, userHistoryController *controller.UserHistoryController) http.Handler	{
	routes := httprouter.New()

	routes.POST("/auth/login", authController.Login)
	routes.POST("/auth/register", authController.Register)
	routes.GET("/user-histories", middleware.AuthMiddleware(userHistoryController.GetUserHistories))
	routes.POST("/user-histories", middleware.AuthMiddleware(userHistoryController.CreateUserHistory))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return c.Handler(routes) 

}