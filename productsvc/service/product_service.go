package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/jeremylo/microsvc/grpc/model/stock"
	"gitlab.com/jeremylo/microsvc/productsvc/domain"
	"gitlab.com/jeremylo/microsvc/productsvc/domain/model"
)

type ProductServiceImpl struct {
	Repository         domain.ProductRepository
	StockServiceClient stock.StockServiceClient
}

func (cs *ProductServiceImpl) FetchProduct(ctx context.Context) []*model.Product {
	productsFromDatabase := cs.Repository.Get(ctx)

	var products []*model.Product
	for _, product := range productsFromDatabase {
		findStockByProduct, err := cs.StockServiceClient.FindStockByProduct(ctx, &stock.Message{Body: int32(product.ID)})
		if err != nil {
			log.WithError(err).Errorln("error when find findStockByProduct product")
		}

		productStock := int(findStockByProduct.Body)
		product.Stock = &productStock
		products = append(products, product)
	}

	return products
}

func (cs *ProductServiceImpl) ShowProduct(ctx context.Context, id *int) *model.Product {
	return cs.Repository.Show(ctx, id)
}
