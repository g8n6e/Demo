package main

type cage struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Capacity string `json:"capacity"`
	Active   bool   `json:"active"`
}
