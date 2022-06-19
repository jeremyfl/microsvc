package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/mocks"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
	"testing"
)

var (
	ctx = context.Background()
	r   = mocks.OrderRepository{Mock: mock.Mock{}}
	m   = mocks.MessageBroker{Mock: mock.Mock{}}

	s = &OrderServiceImpl{
		Repository:    &r,
		MessageBroker: &m,
	}

	orderData = &model.Order{
		ProductID:  1,
		IsCanceled: false,
	}
)

func TestOrderServiceImpl_CancelOrder(t *testing.T) {
	r.On("Update", ctx, mock.Anything, mock.Anything).Return(nil)

	assert.Nil(t, s.CancelOrder(ctx, orderData))
}

func TestOrderServiceImpl_CreateOrder(t *testing.T) {
	r.On("Create", ctx, mock.Anything).Return(orderData, nil)
	m.On("Publish", ctx, "order.created", mock.Anything).Return(nil)

	assert.Nil(t, s.CreateOrder(ctx, orderData))
}
