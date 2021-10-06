package main

import (
	"github.com/vandyahmad24/gin-app/config"
	"github.com/vandyahmad24/gin-app/router"
)

func main() {
	config.InitDB()
	router.Router()
}
