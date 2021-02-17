package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is used to access the database
var DB *gorm.DB

// ConnectDatabase opens the database file and migrates the models
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Excercise{})
	database.AutoMigrate(&Session{})
	database.AutoMigrate(&User{})
	DB = database
}
