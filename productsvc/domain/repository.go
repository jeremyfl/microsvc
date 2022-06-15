package domain

import (
	"context"
	"gitlab.com/jeremylo/microsvc/productsvc/domain/model"
)

type ProductRepository interface {
	Get(ctx context.Context) []*model.Product
	Show(ctx context.Context, id *int) *model.Product
}
