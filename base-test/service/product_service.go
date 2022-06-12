package service

import (
	"context"
	"customer/domain"
	"customer/domain/model"
)

type ProductServiceImpl struct {
	Repository domain.ProductRepository
}

func (cs ProductServiceImpl) FetchProduct(ctx context.Context) *[]model.Product {
	return cs.Repository.Get(ctx)
}