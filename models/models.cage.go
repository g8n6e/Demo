package models

type Cage struct {
	ID       int  `json:"id" gorm:"primaryKey"`
	Capacity int  `json:"capacity"`
	Active   bool `json:"active"`
}
