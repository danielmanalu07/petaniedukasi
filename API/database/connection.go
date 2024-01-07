package database

import (
	"petani_edukasi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/petaniedukasi"), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Admin{})
	connection.AutoMigrate(&models.Edukasi{})
}
