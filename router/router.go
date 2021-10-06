package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vandyahmad24/gin-app/config"
	"github.com/vandyahmad24/gin-app/controller"
	"github.com/vandyahmad24/gin-app/entity/user"
)

func Router() {
	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	userHandler := controller.NewUserHandler(userService)
	// userService.SaveAvatar(21, "main.jpg")

	r := gin.Default()
	api := r.Group("api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailIsUse)
	api.POST("/avatars", userHandler.UploadAvatar)
	r.Run(":9191")
}
