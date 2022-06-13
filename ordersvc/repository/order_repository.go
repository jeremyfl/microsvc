package repository

import (
	"golang.org/x/net/context"
	"ordersvc/domain/model"
	"ordersvc/internal"
)

// OrderRepositoryImpl is the repository structure
type OrderRepositoryImpl struct {
	DB *internal.Database
}

func (cr *OrderRepositoryImpl) Create(ctx context.Context, payload *model.Order) error {
	cr.DB.Create(payload)

	return nil
}