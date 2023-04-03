package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	NameFood    string  `gorm:"not null" json:"name_food"`
	Price       float32 `gorm:"not null" json:"price"`
	Despription string  `gorm:"not null" json:"description"`
	Category    string  `gorm:"not null" json:"category"`
	FoodId      uint    `json:"food_id"`
	ImageUrl    string  `json:"image"`
}
