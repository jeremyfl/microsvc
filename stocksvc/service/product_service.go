package service

import (
	"context"
	"gitlab.com/jeremylo/microsvc/grpc/model/stock"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain/model"
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

func (cs *StockServiceImpl) DecreaseStock(ctx context.Context, productID int) error {
	currentStock := cs.Repository.Show(ctx, &model.Stock{ProductID: productID})
	nextStock := currentStock.Total - 1

	if err := cs.Repository.Update(ctx, currentStock, &model.Stock{Total: nextStock}); err != nil {
		return err
	}

	// simulating the stock service needs to process logic for a while
	//time.Sleep(time.Second * 5)

	return nil
}

func (cs *StockServiceImpl) IncreaseStock(ctx context.Context, productID int) error {
	currentStock := cs.Repository.Show(ctx, &model.Stock{ProductID: productID})
	currentStock.Total += 1

	return cs.Repository.Update(ctx, currentStock, &model.Stock{Total: currentStock.Total})
}
