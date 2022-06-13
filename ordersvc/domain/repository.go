package domain

import (
	"context"
	"ordersvc/domain/model"
)

type OrderRepository interface {
	Create(ctx context.Context, payload *model.Order) error
}
