package domain

import (
	"context"
	"productsvc/domain/model"
)

type Services struct {
	ProductService
}

type ProductService interface {
	FetchProduct(ctx context.Context) []*model.Product
	ShowProduct(ctx context.Context, id *int) *model.Product
}
