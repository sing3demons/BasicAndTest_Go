package config

import (
	"app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}



// SetupDB - setup dabase for sharing to all api
func SetupDB() {


	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})	
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&models.User{})

	db = database
}