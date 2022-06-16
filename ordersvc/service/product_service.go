package service

import (
	"context"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
)

type StockServiceImpl struct {
	Repository domain.OrderRepository
	MessageBroker domain.MessageBroker
}

func (cs *StockServiceImpl) CreateOrder(ctx context.Context, payload *model.Order) error {
	_, span := domain.Tracer.Start(ctx, "CreateOrder")
	defer span.End()

	if err := cs.Repository.Create(ctx, payload); err != nil {
		return err
	}

	if err := cs.MessageBroker.Publish(ctx, "order.created", payload); err != nil {
		return err
	}

	return nil
}

func (cs *StockServiceImpl) CancelOrder(ctx context.Context, payload *model.Order) error {
	_, span := domain.Tracer.Start(ctx, "CreateOrder")
	defer span.End()

	if err := cs.Repository.Update(ctx, payload, &model.Order{IsCanceled: true}); err != nil {
		return err
	}

	if err := cs.MessageBroker.Publish(ctx, "order.canceled", payload); err != nil {
		return err
	}

	return nil
}
