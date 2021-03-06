package models

import (
	"gorm.io/gorm"
)

// Excercise is an activity done during a session
type Excercise struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Weight    string `json:"weight"`
	SessionID uint   `json:"session"`
}
