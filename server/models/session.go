package models

import (
   "gorm.io/gorm"
   "time"
)

type Session struct {
  gorm.Model
  ID            uint   	     `json:"id" gorm:"primary_key"`
  SessionDate   time.Time    `json:"sessionDate"`
  Duration      int64        `json:"duration" gorm:"default:0"`
  Excercises    []Excercise  `json:"excercises"`
}