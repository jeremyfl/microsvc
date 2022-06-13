package domain

import (
	"context"
	"stocksvc/domain/model"
)

type Services struct {
	StockService
}

type StockService interface {
	FetchStock(ctx context.Context) []*model.Stock
	ShowStock(ctx context.Context, productID int) *model.Stock
}
