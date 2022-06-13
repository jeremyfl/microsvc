package domain

import (
	"context"
	"productsvc/domain/model"
)

type ProductRepository interface {
	Get(ctx context.Context) []*model.Product
}
