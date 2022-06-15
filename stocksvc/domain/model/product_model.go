package model

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ProductID int
	Total     int
}