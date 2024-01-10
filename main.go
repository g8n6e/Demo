package main

import (
	"prizepicks/jurassicpark/models"
	"prizepicks/jurassicpark/routes"
)

func main() {
	models.ConnectDatabase()
	routes.Run()
}
