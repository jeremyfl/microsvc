package service

import (
	"context"
	"github.com/segmentio/kafka-go"
	"ordersvc/domain"
	"ordersvc/domain/model"
)

type StockServiceImpl struct {
	Repository domain.OrderRepository
	Publisher  *kafka.Writer
}

func (cs *StockServiceImpl) CreateOrder(ctx context.Context, payload *model.Order) error {
	if err := cs.Repository.Create(ctx, payload); err != nil {
		return err
	}

	m := kafka.Message{
		Value: []byte("Hello World!"),
	}

	if err := cs.Publisher.WriteMessages(ctx, m); err != nil {
		return err
	}

	return nil
}
