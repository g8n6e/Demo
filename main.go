package main

import (
	"log"
	"prizepicks/jurassicpark/routes"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := connectToSQLite()
	if err != nil {
		log.Fatal(err)
	}

	// Perform database migration
	err = db.AutoMigrate(&cage{})
	if err != nil {
		log.Fatal(err)
	}
	routes.Run()
}

func connectToSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
