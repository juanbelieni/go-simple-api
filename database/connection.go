package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/juanbelieni/go-simple-api/models"
)

func getConnection() *gorm.DB {
	c, err := gorm.Open("sqlite3", "database.sqlite")
	if err != nil {
		panic("db: Failed to connect to database")
	}

	c.AutoMigrate(&models.User{})

	return c
}

var Connection = getConnection()
