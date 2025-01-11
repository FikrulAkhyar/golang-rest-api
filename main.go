package main

import (
	"os"

	"github.com/FikrulAkhyar/golang-rest-api/controllers/userController"
	"github.com/FikrulAkhyar/golang-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/users", userController.Index)
	r.GET("/api/users/:id", userController.Show)
	r.POST("/api/users", userController.Store)
	r.PUT("/api/users/:id", userController.Update)
	r.DELETE("/api/users/:id", userController.Delete)

	r.Run(":" + port)
}
