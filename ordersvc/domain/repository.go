package domain

import (
	"context"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
)

type OrderRepository interface {
	Create(ctx context.Context, payload *model.Order) error
	Update(ctx context.Context, filter, payload *model.Order) error
}
