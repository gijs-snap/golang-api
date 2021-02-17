package main

import (
  "github.com/gin-gonic/gin"
  "gymgo.com/m/models"
  "github.com/joho/godotenv"
  "gymgo.com/m/controllers"
  "log"
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
  r.GET("/excercises", controllers.FindExcercises)
  r.POST("/session/create", controllers.CreateSession)
  r.Run(":2000")
}