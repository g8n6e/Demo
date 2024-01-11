package handlers

import (
	"errors"
	"prizepicks/jurassicpark/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dinosaurRequest struct {
	Name     string `json:"name"`
	SpecieID int    `json:"specieId"`
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
	if input.SpecieID == 0 {
		return models.Dinosaur{}, errors.New("cannot add a dinosaur without a specieid")
	}
	if input.Name == "" {
		return models.Dinosaur{}, errors.New("dinosaurs name must be provided")
	}
	if input.CageID == 0 {
		//add logic to find available space in existing cages
		cage, err := addCage(cageRequest{})
		if err != nil {
			return dinosaur, err
		}
		input.CageID = cage.ID
	} else {
		cage, err := getCageById(input.CageID)
		if err != nil {
			return models.Dinosaur{}, err
		}
		if !cage.Active {
			return models.Dinosaur{}, errors.New("this cage is down please find another cage")
		}
		dinosaurs, err := getDinosaursByCageId(input.CageID)
		if err != nil {
			return models.Dinosaur{}, err
		}
		if cage.Capacity <= len(dinosaurs) {
			return models.Dinosaur{}, errors.New("this cage cannot hold another dinosaur please find another cage")
		}
		if len(dinosaurs) != 0 {
			specie, err := getSpecieById(input.SpecieID)
			if err != nil {
				return models.Dinosaur{}, err
			}
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
	if dinosaur.CageID != input.CageID {
		cageId := dinosaur.CageID
		if input.CageID != 0 {
			cageId = input.CageID
		}
		specieId := dinosaur.SpecieID
		if input.SpecieID != 0 {
			specieId = input.SpecieID
		}
		cage, err := getCageById(cageId)
		if err != nil {
			return models.Dinosaur{}, err
		}
		if !cage.Active {
			return models.Dinosaur{}, errors.New("this cage is down please find another cage")
		}
		dinosaurs, err := getDinosaursByCageId(cageId)
		if err != nil {
			return models.Dinosaur{}, err
		}
		if cage.Capacity <= len(dinosaurs) {
			return models.Dinosaur{}, errors.New("this cage cannot hold another dinosaur please find another cage")
		}
		if len(dinosaurs) != 0 {
			specie, err := getSpecieById(specieId)
			if err != nil {
				return models.Dinosaur{}, err
			}
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
	err = models.DB.Model(&dinosaur).Updates(input).Error
	return dinosaur, err
}
