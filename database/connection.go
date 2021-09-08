package database

import (
	"github.com/ronanzindev/goauthfiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:morellianogm12321@tcp(127.0.0.1:3306)/goauthfiber?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = connection

	connection.AutoMigrate(&models.User{})
}
