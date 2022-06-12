package repository

import (
	"customer/domain/model"
	"customer/internal"
	"golang.org/x/net/context"
)

// ProductRepositoryImpl is the repository structure
type ProductRepositoryImpl struct {
	DB *internal.Database
}

func (cr ProductRepositoryImpl) Get(ctx context.Context) *[]model.Product {
	var products []model.Product

	cr.DB.Find(&products)

	return &products
}
