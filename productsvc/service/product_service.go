package service

import (
	"context"
	"fmt"
	"gitlab.com/jeremylo/microsvc/grpc/model/stock"
	"gitlab.com/jeremylo/microsvc/productsvc/domain"
	"gitlab.com/jeremylo/microsvc/productsvc/domain/model"
)

type ProductServiceImpl struct {
	Repository         domain.ProductRepository
	StockServiceClient stock.StockServiceClient
}

func (cs *ProductServiceImpl) FetchProduct(ctx context.Context) []*model.Product {
	x, err := cs.StockServiceClient.FindStockByProduct(ctx, &stock.Message{Body: 123})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(x.Body)

	return cs.Repository.Get(ctx)
}

func (cs *ProductServiceImpl) ShowProduct(ctx context.Context, id *int) *model.Product {
	return cs.Repository.Show(ctx, id)
}