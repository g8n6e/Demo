package models

type Dinosaur struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	SpecieID int    `json:"specieId"`
	CageID   int    `json:"cageId"`
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
