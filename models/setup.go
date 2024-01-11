package models

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dbLocation string) {

	database, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	if err = database.AutoMigrate(&Specie{}); err == nil && database.Migrator().HasTable(&Specie{}) {
		if err := database.First(&Specie{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			species := []Specie{
				{Name: "Brachiosaurus", Diet: 0},
				{Name: "Stegosaurus", Diet: 0},
				{Name: "Ankylosaurus", Diet: 0},
				{Name: "Triceratops", Diet: 0},
				{Name: "Tyrannosaurus", Diet: 1},
				{Name: "Velociraptor", Diet: 1},
				{Name: "Spinosaurus", Diet: 1},
				{Name: "Megalosaurus", Diet: 1},
			}
			result := database.Create(species)
			if result.Error != nil {
				return
			}
		}
	}
	if err != nil {
		return
	}
	err = database.AutoMigrate(&Cage{}, &Dinosaur{})
	if err != nil {
		return
	}

	DB = database
}
