package handlers

import (
	"prizepicks/jurassicpark/models"

	"github.com/gin-gonic/gin"
)

type CreateCageRequest struct {
	Capacity int  `json:"capacity" binding:"required"`
	Active   bool `json:"active" binding:"required"`
}

func GetCages() []models.Cage {
	var cages []models.Cage
	models.DB.Find(&cages)
	return cages
}

func AddCage(c *gin.Context) (cage models.Cage, err error) {
	var input CreateCageRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return cage, err
	}

	// Create book
	newCage := models.Cage{Capacity: input.Capacity, Active: input.Active}
	models.DB.Create(&newCage)
	return newCage, nil
}
