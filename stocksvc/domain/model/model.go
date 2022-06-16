package model

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ProductID int
	Total     int
}

type OrderPayload struct {
	ID            int  `json:"id"`
	ProductID     int  `json:"product_id" xml:"product_id" form:"product_id"`
	TotalPrice    int  `json:"total_price" xml:"total_price" form:"total_price"`
	TotalQuantity int  `json:"total_quantity" xml:"total_quantity" form:"total_quantity"`
	IsCanceled    bool `json:"is_canceled"`
}
