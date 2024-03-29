package repository

import (
	"gitlab.com/jeremylo/microsvc/stocksvc/domain/model"
	"gitlab.com/jeremylo/microsvc/stocksvc/internal"
	"golang.org/x/net/context"
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

func (cr *StockRepositoryImpl) Update(ctx context.Context, filter *model.Stock, payload *model.Stock) error {
	cr.DB.Model(&filter).Where("product_id", filter.ProductID).Updates(payload)

	return nil
}
