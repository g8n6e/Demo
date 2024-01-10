package handlers

import (
	"prizepicks/jurassicpark/models"

	"github.com/gin-gonic/gin"
)

type dinosaurRequest struct {
	Name      string `json:"name"`
	SpeciesID int    `json:"speciesId"`
	CageID    int    `json:"cageId"`
}

func GetDinosaurs() []models.Dinosaur {
	var dinosaurs []models.Dinosaur
	models.DB.Find(&dinosaurs)
	return dinosaurs
}

func AddDinosaur(c *gin.Context) (dinosaur models.Dinosaur, err error) {
	var input dinosaurRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return dinosaur, err
	}
	newDinosaur := models.Dinosaur{Name: input.Name, CageID: input.CageID, SpeciesID: input.SpeciesID}
	err = models.DB.Create(&newDinosaur).Error
	return newDinosaur, err
}

func GetDinosaurById(c *gin.Context) (models.Dinosaur, error) {
	var dinosaur models.Dinosaur
	err := models.DB.Where("id = ?", c.Param("id")).First(&dinosaur).Error
	return dinosaur, err
}

func UpdateDinosaur(c *gin.Context) (dinosaur models.Dinosaur, err error) {
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dinosaur).Error; err != nil {
		return dinosaur, err
	}
	var input dinosaurRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return dinosaur, err
	}
	err = models.DB.Model(&dinosaur).Updates(input).Error
	return dinosaur, err
}
