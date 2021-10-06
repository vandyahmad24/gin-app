package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vandyahmad24/gin-app/controller"
	"github.com/vandyahmad24/gin-app/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3310)/campaign?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := controller.NewUserHandler(userService)

	r := gin.Default()
	api := r.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	r.Run(":9191")
	// userRepository.Save(user)
	// fmt.Println("koneksi sukses")
	// var users []user.User
	// db.Find(&users)
	// for _, v := range users {
	// 	fmt.Println(v.Name)
	// }

	// r := gin.Default()
	// r.GET("/", HandlerUser)
	// r.Run(":9191")

}
