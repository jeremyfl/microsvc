package service

import (
	"context"
	"productsvc/domain"
	"productsvc/domain/model"
	"productsvc/stock"
)

type ProductServiceImpl struct {
	Repository domain.ProductRepository
	StockServiceClient stock.StockServiceClient
}

func (cs *ProductServiceImpl) FetchProduct(ctx context.Context) []*model.Product {
	return cs.Repository.Get(ctx)
}

func (cs *ProductServiceImpl) ShowProduct(ctx context.Context, id *int) *model.Product {
	return cs.Repository.Show(ctx, id)
}