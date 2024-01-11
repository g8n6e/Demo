package handlers

import (
	"errors"
	"prizepicks/jurassicpark/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dinosaurRequest struct {
	Name     string `json:"name"`
	SpecieID int    `json:"speciesId"`
	CageID   int    `json:"cageId"`
}

func GetDinosaurs() ([]models.Dinosaur, error) {
	var dinosaurs []models.Dinosaur
	err := models.DB.Find(&dinosaurs).Error
	return dinosaurs, err
}

func GetDinosaursByCageId(c *gin.Context) ([]models.Dinosaur, error) {
	cageId, conversionErr := strconv.Atoi(c.Param("id"))
	if conversionErr != nil {
		return []models.Dinosaur{}, conversionErr
	}
	dinosaurs, err := getDinosaursByCageId(cageId)
	return dinosaurs, err
}

func getDinosaursByCageId(cageId int) ([]models.Dinosaur, error) {
	var dinosaurs []models.Dinosaur
	err := models.DB.Where("cage_id = ?", cageId).Find(&dinosaurs).Error
	if err != nil {
		return []models.Dinosaur{}, err
	}
	return dinosaurs, err
}

func AddDinosaur(c *gin.Context) (dinosaur models.Dinosaur, err error) {
	var input dinosaurRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		return dinosaur, err
	}
	if input.CageID == 0 {
		//add logic to find available space in existing cages
		cage, err := addCage(cageRequest{})
		if err != nil {
			return dinosaur, err
		}
		input.CageID = cage.ID
	} else {
		specie, err := getSpecieById(input.SpecieID)
		if err != nil {
			return models.Dinosaur{}, err
		}
		dinosaurs, err := getDinosaursByCageId(input.CageID)
		if err != nil {
			return models.Dinosaur{}, err
		}
		if len(dinosaurs) != 0 {
			cageSpecie, err := getSpecieById(dinosaurs[0].SpecieID)
			if err != nil {
				return models.Dinosaur{}, err
			}
			if specie.Diet == 1 && cageSpecie.ID != specie.ID {
				return models.Dinosaur{}, errors.New("Cannot add " + specie.Name + " to cage with " + cageSpecie.Name)
			} else if specie.Diet == 0 && cageSpecie.Diet == 1 {
				return models.Dinosaur{}, errors.New("Cannot add " + specie.Name + " to cage with carnivore")
			}
		}
	}
	newDinosaur := models.Dinosaur{Name: input.Name, CageID: input.CageID, SpecieID: input.SpecieID}
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
	specie, err := getSpecieById(dinosaur.SpecieID)
	if err != nil {
		return models.Dinosaur{}, err
	}
	dinosaurs, err := getDinosaursByCageId(dinosaur.CageID)
	if err != nil {
		return models.Dinosaur{}, err
	}
	if dinosaurs != nil {
		cageSpecie, err := getSpecieById(dinosaurs[0].SpecieID)
		if err != nil {
			return models.Dinosaur{}, err
		}
		if specie.Diet == 1 && cageSpecie.ID != specie.ID {
			return models.Dinosaur{}, errors.New("Cannot add " + specie.Name + " to cage with " + cageSpecie.Name)
		} else if specie.Diet == 0 && cageSpecie.Diet == 1 {
			return models.Dinosaur{}, errors.New("Cannot add " + specie.Name + " to cage with carnivore")
		}
	}
	err = models.DB.Model(&dinosaur).Updates(input).Error
	return dinosaur, err
}
