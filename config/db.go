package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connDB := "root:@tcp(127.0.0.1:3310)/campaign?charset=utf8mb4&parseTime=True&loc=Local"

	var e error
	DB, e = gorm.Open(mysql.Open(connDB), &gorm.Config{})
	if e != nil {
		panic(e)
	}

}
