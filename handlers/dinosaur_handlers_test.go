package handlers

import (
	"prizepicks/jurassicpark/models"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCageSegregation(t *testing.T) {
	models.ConnectDatabase("../demo.db")
	cages := GetCages()
	for _, cage := range cages {
		dinosaurs, err := getDinosaursByCageId(cage.ID)
		if err != nil {
			assert.Fail(t, "Error Getting Dinosaurs for cageid: "+strconv.Itoa(cage.ID))
		}
		if dinosaurs != nil {
			initalDinosaur := dinosaurs[0]
			initalDinosaurSpecie, err := getSpecieById(initalDinosaur.SpecieID)
			if err != nil {
				assert.Fail(t, "Error getting dinosaur specie for dino: "+strconv.Itoa(initalDinosaur.ID))
			}
			for _, dinosaur := range dinosaurs {
				if initalDinosaurSpecie.Diet == 1 {
					assert.Equal(t, initalDinosaurSpecie.ID, dinosaur.SpecieID)
				} else {
					specie, err := getSpecieById(dinosaur.SpecieID)
					if err != nil {
						assert.Fail(t, "Error getting dinosaur specie for dino: "+strconv.Itoa(dinosaur.ID))
					}
					assert.Equal(t, specie.Diet, models.Herbivore)
				}
			}
		}
	}
}
