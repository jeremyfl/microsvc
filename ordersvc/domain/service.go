package domain

import (
	"context"
	"ordersvc/domain/model"
)

type Services struct {
	StockService
}

type StockService interface {
	CreateOrder(ctx context.Context, payload *model.Order) error
}
