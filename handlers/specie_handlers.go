package handlers

import (
	"errors"
	"prizepicks/jurassicpark/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type specieRequest struct {
	Name string      `json:"name"`
	Diet models.Diet `json:"diet"`
}

func GetSpecies() ([]models.Specie, error) {
	var species []models.Specie
	err := models.DB.Find(&species).Error
	return species, err
}

func AddSpecie(c *gin.Context) (specie models.Specie, err error) {
	var input specieRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return specie, err
	}
	if input.Name == "" {
		return models.Specie{}, errors.New("specie name must be provided")
	}
	newSpecie := models.Specie{Name: input.Name, Diet: input.Diet}
	err = models.DB.Create(&newSpecie).Error
	return newSpecie, err
}

func GetSpecieById(c *gin.Context) (models.Specie, error) {
	specieId, conversionErr := strconv.Atoi(c.Param("id"))
	if conversionErr != nil {
		return models.Specie{}, conversionErr
	}
	specie, err := getSpecieById(specieId)
	return specie, err
}

func getSpecieById(specieId int) (models.Specie, error) {
	var specie models.Specie
	err := models.DB.Where("id = ?", specieId).First(&specie).Error
	return specie, err
}

func UpdateSpecie(c *gin.Context) (specie models.Specie, err error) {
	if err := models.DB.Where("id = ?", c.Param("id")).First(&specie).Error; err != nil {
		return specie, err
	}
	var input specieRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return specie, err
	}
	err = models.DB.Model(&specie).Updates(input).Error
	return specie, err
}
