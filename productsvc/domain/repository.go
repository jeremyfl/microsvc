package domain

import (
	"context"
	"productsvc/domain/model"
)

type ProductRepository interface {
	Get(ctx context.Context) []*model.Product
	Show(ctx context.Context, id *int) *model.Product
}
