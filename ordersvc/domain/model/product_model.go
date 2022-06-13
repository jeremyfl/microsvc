package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID int
	TotalPrice int
	TotalQuantity int
}
