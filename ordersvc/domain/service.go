package domain

import (
	"context"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
)

type Services struct {
	OrderService
}

type OrderService interface {
	CreateOrder(ctx context.Context, payload *model.Order) error
	CancelOrder(ctx context.Context, payload *model.Order) error
}
