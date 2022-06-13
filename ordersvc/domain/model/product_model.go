package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID     int `json:"product_id" xml:"product_id" form:"product_id"`
	TotalPrice    int `json:"total_price" xml:"total_price" form:"total_price"`
	TotalQuantity int `json:"total_quantity" xml:"total_quantity" form:"total_quantity"`
}
