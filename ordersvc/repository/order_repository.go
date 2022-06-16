package repository

import (
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
	"gitlab.com/jeremylo/microsvc/ordersvc/internal"
	"golang.org/x/net/context"
)

// OrderRepositoryImpl is the repository structure
type OrderRepositoryImpl struct {
	DB *internal.Database
}

func (cr *OrderRepositoryImpl) Create(ctx context.Context, payload *model.Order) (*model.Order, error) {
	_, span := domain.Tracer.Start(ctx, "Create")
	defer span.End()

	cr.DB.Create(payload)

	return payload, nil
}

func (cr *OrderRepositoryImpl) Update(ctx context.Context, filter, payload *model.Order) error {
	_, span := domain.Tracer.Start(ctx, "Create")
	defer span.End()

	cr.DB.Model(&filter).Where("id", filter.ID).Updates(payload)

	return nil
}
