package repository

import (
	"golang.org/x/net/context"
	"productsvc/domain/model"
	"productsvc/internal"
)

// ProductRepositoryImpl is the repository structure
type ProductRepositoryImpl struct {
	DB *internal.Database
}

func (cr *ProductRepositoryImpl) Get(ctx context.Context) []*model.Product {
	var products []*model.Product

	cr.DB.Find(&products)

	return products
}

func (cr *ProductRepositoryImpl) Show(ctx context.Context, id *int) *model.Product {
	var product model.Product

	cr.DB.First(&product, id)

	return &product
}
