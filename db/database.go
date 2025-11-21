package db

import (
	"log"
	"os"

	"github.com/ingenziart/myapp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connecting to db
var DB *gorm.DB

func ConnectingDb() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("fail to connect to db ", err)

	}
	println("succesfully connected")

	//auto migrate
	err = DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("failed to migrate ", err)
	}
	println("succesfully migrated")

}
