package service

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain/model"
)

type StockServiceImpl struct {
	Repository domain.OrderRepository
	Publisher  *kafka.Writer
}

func (cs *StockServiceImpl) CreateOrder(ctx context.Context, payload *model.Order) error {
	if err := cs.Repository.Create(ctx, payload); err != nil {
		return err
	}

	p, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	m := kafka.Message{
		Value: p,
	}

	if err := cs.Publisher.WriteMessages(ctx, m); err != nil {
		return err
	}

	return nil
}
