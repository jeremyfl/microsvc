package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/jeremylo/microsvc/grpc/model/stock"
	"gitlab.com/jeremylo/microsvc/productsvc/domain/mocks"
	"gitlab.com/jeremylo/microsvc/productsvc/domain/model"
	ms "gitlab.com/jeremylo/microsvc/productsvc/service/mocks"
	"testing"
)

var (
	stockClient = ms.StockServiceClientMock{Mock: mock.Mock{}}
	repo = mocks.ProductRepository{Mock: mock.Mock{}}
	p    = &ProductServiceImpl{&repo, &stockClient}

	productName = "Apple"
	productData = &model.Product{
		Name: &productName,
	}
)
func TestProductServiceImpl_FetchProduct(t *testing.T) {
	ctx := context.Background()

	var expectedResult []*model.Product
	expectedResult = append(expectedResult, productData)

	repo.On("Get", ctx).Return(expectedResult, nil)
	stockClient.On("FindStockByProduct", ctx, mock.Anything).Return(&stock.Message{Body: 0}, nil)

	products := p.FetchProduct(ctx)

	assert.Equal(t, expectedResult, products)
}

func TestProductServiceImpl_ShowProduct(t *testing.T) {
	ctx := context.Background()

	productId := 1
	repo.On("Show", ctx, &productId).Return(productData, nil)
	product := p.ShowProduct(ctx, &productId)

	assert.Equal(t, productData, product)
}
