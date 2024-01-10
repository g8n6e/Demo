package main

import (
	"prizepicks/jurassicpark/models"
	"prizepicks/jurassicpark/routes"
)

func main() {
	models.ConnectDatabase("demo.db")
	routes.Run()
}
