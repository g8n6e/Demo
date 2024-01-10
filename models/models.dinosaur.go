package models

type Dinosaur struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	SpeciesID int    `json:"speciesId"`
	Species   Specie `json:"species"`
	CageID    int    `json:"cageId"`
	Cage      Cage   `json:"cage"`
}

type Specie struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Diet Diet   `json:"diet"`
}

type Diet int

const (
	Herbivore Diet = iota
	Carnivore
)
