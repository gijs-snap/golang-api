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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, origin, Cache-Control, X-Requested-With, Set-Cookie, X-HTTP-Method-Override")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	models.ConnectDatabase()
	r.POST("/user/login", controllers.LoginUser)
	r.POST("/user/register", controllers.RegisterUser)
	r.GET("/excercises", controllers.FindExcercises)
	r.POST("/session/create", controllers.CreateSession)
	r.Run(":2000")
}
