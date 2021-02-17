package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gymgo.com/m/controllers"
	"gymgo.com/m/models"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.POST("/user/login", controllers.LoginUser)
	r.POST("/user/register", controllers.RegisterUser)
	r.GET("/excercises", controllers.FindExcercises)
	r.POST("/session/create", controllers.CreateSession)
	r.Run(":2000")
}
