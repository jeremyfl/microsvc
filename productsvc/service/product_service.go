package service

import (
	"context"
	"productsvc/domain"
	"productsvc/domain/model"
)

type ProductServiceImpl struct {
	Repository domain.ProductRepository
}

func (cs ProductServiceImpl) FetchProduct(ctx context.Context) []*model.Product {
	return cs.Repository.Get(ctx)
}
