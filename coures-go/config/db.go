package config

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

  var db *gorm.DB

func InitDB(){
	var err error
	db, err = gorm.Open("sqlite3","./gorm.db")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(gin.Mode() = gin.DebugMode)
}

func GetDB() *gorm.DB {
	return db
}
func CloseDB(){
	db.Close()
}
