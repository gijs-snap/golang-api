package models

import (
	"gorm.io/gorm"
)

// User is an entity using the API
type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
