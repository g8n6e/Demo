package models

type Cage struct {
	ID       int  `json:"id" gorm:"primaryKey"`
	Capacity int  `json:"capacity" gorm:"default:1"`
	Active   bool `json:"active" gorm:"default:true"`
}
