package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sharangj/weather_monster/config"

	// Gorm docs says there are cases where the api may need a seperate variable to store imports when there are multiple DB'S
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Connect : Initialises the DB Connection
func Connect() *gorm.DB {
	c := config.GetConfig()
	db, err := gorm.Open("postgres", c.GetString("dbString"))
	if err != nil {
		panic(err)
	}
	db.LogMode(false) // Set to true for debugging
	return db
}
