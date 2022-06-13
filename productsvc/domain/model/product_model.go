package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        *string  `json:"name"`
	Price       *float64 `json:"price"`
	Description *string  `json:"description"`
	Rating      *float64 `json:"rating"`
	Category    *string  `json:"category"`
	Brand       *string  `json:"brand"`
	Thumbnail   *string  `json:"thumbnail"`
}
