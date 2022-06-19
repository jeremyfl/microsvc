package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain/model"
	"time"
)

type StockServiceImpl struct {
	Repository         domain.StockRepository
	MessageBroker      domain.MessageBroker
}

func (cs *StockServiceImpl) FetchStock(ctx context.Context) []*model.Stock {
	return cs.Repository.Get(ctx)
}

func (cs *StockServiceImpl) ShowStock(ctx context.Context, productID int) *model.Stock {
	payload := &model.Stock{ProductID: productID}

	return cs.Repository.Show(ctx, payload)
}

func (cs *StockServiceImpl) DecreaseStock(ctx context.Context, payload *model.OrderPayload) error {
	currentStock := cs.Repository.Show(ctx, &model.Stock{ProductID: payload.ProductID})

	log.WithField("current_stock_total", currentStock.Total).Infoln("decreasing ..")

	if currentStock.Total < payload.TotalQuantity {
		if err := cs.MessageBroker.Publish(ctx, "stock.exceeded-amount", &payload); err != nil {
			return err
		}

		log.WithField("product_id", payload.ProductID).Infoln("stock exceeded, reverting..")

		return nil
	}

	nextStock := currentStock.Total - payload.TotalQuantity

	if err := cs.Repository.Update(ctx, currentStock, &model.Stock{Total: nextStock}); err != nil {
		return err
	}

	//simulating the stock service needs to process logic for a while
	time.Sleep(time.Second * 5)

	log.WithField("product_id", payload.ProductID).Infoln("stock decreased")

	return nil
}

func (cs *StockServiceImpl) IncreaseStock(ctx context.Context, payload *model.OrderPayload) error {
	currentStock := cs.Repository.Show(ctx, &model.Stock{ProductID: payload.ProductID})
	currentStock.Total += payload.TotalQuantity

	return cs.Repository.Update(ctx, currentStock, &model.Stock{Total: currentStock.Total})
}
