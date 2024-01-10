package handlers

import (
	"prizepicks/jurassicpark/models"

	"github.com/gin-gonic/gin"
)

type cageRequest struct {
	Capacity int  `json:"capacity"`
	Active   bool `json:"active"`
}

func GetCages() []models.Cage {
	var cages []models.Cage
	models.DB.Find(&cages)
	return cages
}

func AddCage(c *gin.Context) (cage models.Cage, err error) {
	var input cageRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return cage, err
	}
	newCage := models.Cage{Capacity: input.Capacity, Active: input.Active}
	err = models.DB.Create(&newCage).Error
	return newCage, err
}

func GetCageById(c *gin.Context) (models.Cage, error) {
	var cage models.Cage
	err := models.DB.Where("id = ?", c.Param("id")).First(&cage).Error
	return cage, err
}

func UpdateCage(c *gin.Context) (cage models.Cage, err error) {
	if err := models.DB.Where("id = ?", c.Param("id")).First(&cage).Error; err != nil {
		return cage, err
	}
	var input cageRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return cage, err
	}
	err = models.DB.Model(&cage).Updates(input).Error
	return cage, err
}
