package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dbLocation string) {

	database, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Cage{}, &Dinosaur{}, &Specie{})
	if err != nil {
		return
	}

	DB = database
}
