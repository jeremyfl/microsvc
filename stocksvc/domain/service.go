package domain

import (
	"context"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain/model"
)

type Services struct {
	StockService
}

type StockService interface {
	FetchStock(ctx context.Context) []*model.Stock
	ShowStock(ctx context.Context, productID int) *model.Stock
	DecreaseStock(ctx context.Context, productID int) error
	IncreaseStock(ctx context.Context, productID int) error
}
