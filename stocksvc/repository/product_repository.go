package repository

import (
	"golang.org/x/net/context"
	"stocksvc/domain/model"
	"stocksvc/internal"
)

// StockRepositoryImpl is the repository structure
type StockRepositoryImpl struct {
	DB *internal.Database
}

func (cr *StockRepositoryImpl) Get(ctx context.Context) []*model.Stock {
	var stocks []*model.Stock

	cr.DB.Find(&stocks)

	return stocks
}

func (cr *StockRepositoryImpl) Show(ctx context.Context, payload *model.Stock) *model.Stock {
	var stock *model.Stock

	cr.DB.Where(payload).Find(&stock)

	return stock
}
