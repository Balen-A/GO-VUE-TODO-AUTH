package database

import (
	"Go-Backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/GO"), &gorm.Config{})
	if err != nil {
		panic("unable to connect")
	}

	DB = connection
	connection.AutoMigrate(models.User{})
	connection.AutoMigrate(models.Task{})
}
