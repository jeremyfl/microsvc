package domain

import (
	"context"
	"stocksvc/domain/model"
)

type StockRepository interface {
	Get(ctx context.Context) []*model.Stock
	Show(ctx context.Context, payload *model.Stock) *model.Stock
	Update(ctx context.Context, filter *model.Stock, payload *model.Stock) error
}
