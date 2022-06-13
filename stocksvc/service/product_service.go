package service

import (
	"context"
	"stocksvc/domain"
	"stocksvc/domain/model"
	"stocksvc/stock"
)

type StockServiceImpl struct {
	Repository         domain.StockRepository
	StockServiceClient stock.StockServiceClient
}

func (cs *StockServiceImpl) FetchStock(ctx context.Context) []*model.Stock {
	return cs.Repository.Get(ctx)
}

func (cs *StockServiceImpl) ShowStock(ctx context.Context, productID int) *model.Stock {
	payload := &model.Stock{ProductID: productID}

	return cs.Repository.Show(ctx, payload)
}
