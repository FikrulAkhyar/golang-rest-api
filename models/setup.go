package models

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	database, err := gorm.Open(mysql.Open(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306" + ")/" + dbName))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})

	DB = database
}
