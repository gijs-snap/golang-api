package models

import (
	"time"

	"gorm.io/gorm"
)

// Session represents a workout session a user had
type Session struct {
	gorm.Model
	ID          uint        `json:"id" gorm:"primary_key"`
	SessionDate time.Time   `json:"sessionDate"`
	Duration    int64       `json:"duration" gorm:"default:0"`
	Excercises  []Excercise `json:"excercises"`
}
