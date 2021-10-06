package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vandyahmad24/gin-app/config"
	"github.com/vandyahmad24/gin-app/controller"
	"github.com/vandyahmad24/gin-app/user"
)

func Router() {
	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	userHandler := controller.NewUserHandler(userService)

	r := gin.Default()
	api := r.Group("api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	r.Run(":9191")
}
