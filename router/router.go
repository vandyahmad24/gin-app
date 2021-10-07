package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vandyahmad24/gin-app/config"
	"github.com/vandyahmad24/gin-app/controller"
	"github.com/vandyahmad24/gin-app/entity/user"
	"github.com/vandyahmad24/gin-app/lib"
)

func Router() {
	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	authService := lib.NewService()
	userHandler := controller.NewUserHandler(userService, authService)

	// userService.SaveAvatar(21, "main.jpg")

	r := gin.Default()
	api := r.Group("api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailIsUse)
	api.POST("/avatars", userHandler.UploadAvatar)
	r.Run(":9191")
}

// func authMiddleware(c *gin.Context) {
// 	authHeader := c.GetHeader("Authorization")
// 	if !strings.Contains(authHeader, "Bearer") {
// 		// ressponseMessage := gin.H{}
// 		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 		return
// 	}
// 	tokenString := ""
// 	arrayToken := strings.Split(authHeader, " ")
// 	if len(arrayToken) == 2 {
// 		tokenString = arrayToken[1]
// 	}

// 	token, err :=

// }
