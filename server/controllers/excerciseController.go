package controllers

import (
	"github.com/gin-gonic/gin"
	"gymgo.com/m/models"
)

// FindExcercises retrieves all excercises from DB
func FindExcercises(c *gin.Context) {
	verified := VerifyToken(c)
	if !verified {
		c.JSON(404, "Failed to verify")
		return
	}

	var excercises []models.Excercise
	models.DB.Find(&excercises)

	c.JSON(200, excercises)
}
