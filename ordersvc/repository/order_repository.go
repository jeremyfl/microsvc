package repository

import (
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
	"gitlab.com/jeremylo/microsvc/ordersvc/internal"
	"golang.org/x/net/context"
)

// OrderRepositoryImpl is the repository structure
type OrderRepositoryImpl struct {
	DB *internal.Database
}

func (cr *OrderRepositoryImpl) Create(ctx context.Context, payload *model.Order) error {
	cr.DB.Create(payload)

	return nil
}
