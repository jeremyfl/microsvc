package domain

import (
	"context"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
)

type Services struct {
	StockService
}

type StockService interface {
	CreateOrder(ctx context.Context, payload *model.Order) error
}
