package service

import (
	"context"
	"fmt"
	"productsvc/domain"
	"productsvc/domain/model"
	"productsvc/stock"
)

type ProductServiceImpl struct {
	Repository domain.ProductRepository
	StockServiceClient stock.StockServiceClient
}

func (cs ProductServiceImpl) FetchProduct(ctx context.Context) []*model.Product {
	payload := &stock.Message{
		Body: "1",
	}

	fmt.Println(cs.StockServiceClient.FindStockByProduct(ctx, payload))

	return cs.Repository.Get(ctx)
}
