package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain/mocks"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain/model"
	"testing"
)

var (
	ctx = context.Background()
	r   = mocks.StockRepository{Mock: mock.Mock{}}
	m   = mocks.MessageBroker{Mock: mock.Mock{}}

	s = &StockServiceImpl{
		Repository:    &r,
		MessageBroker: &m,
	}

	stockData = &model.Stock{
		ProductID: 1,
		Total:     10,
	}
)

func TestStockServiceImpl_DecreaseStock(t *testing.T) {
	p := &model.OrderPayload{ID: 1}

	r.On("Show", ctx, mock.Anything).Return(stockData, nil).Once()
	r.On("Update", ctx, mock.Anything, mock.Anything).Return(nil).Once()
	m.On("Publish", ctx, "stock.exceeded-amount", mock.Anything, p).Return(nil)

	assert.Nil(t, s.DecreaseStock(ctx, p))
}

func TestStockServiceImpl_FetchStock(t *testing.T) {
	expectedResult := []*model.Stock{stockData}

	r.On("Get", ctx).Return(expectedResult, nil)

	assert.Equal(t, expectedResult, s.FetchStock(ctx))
}

func TestStockServiceImpl_IncreaseStock(t *testing.T) {
	t.SkipNow()
}

func TestStockServiceImpl_ShowStock(t *testing.T) {
	r.On("Show", ctx, &model.Stock{ProductID: 1}).Return(stockData, nil)

	assert.Equal(t, stockData, s.ShowStock(ctx, 1))
}
