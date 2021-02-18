package controllers

import (
	"github.com/gijs-snap/golang-api/models"
	"github.com/gin-gonic/gin"
)

// CreateSession creates a new session record in db
func CreateSession(c *gin.Context) {
	var session models.Session
	c.Bind(&session)
	models.DB.Create(&session)
	c.JSON(201, gin.H{"success": session})

	// c.JSON(422, gin.H{"error": "Fields are empty"})

}
