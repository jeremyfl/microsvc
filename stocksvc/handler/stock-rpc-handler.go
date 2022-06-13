package handler

import (
	"context"
	"fmt"
	"stocksvc/domain"
	"stocksvc/stock"
)

type Handler struct {
	Services domain.Services
}

func (s *Handler) FindStockByProduct(ctx context.Context, message *stock.Message) (*stock.Message, error) {
	return &stock.Message{Body: 123}, nil
}

func (s *Handler) FindEmptyStock(ctx context.Context, message *stock.Message) (*stock.Message, error) {
	fmt.Println(message)

	return &stock.Message{Body: 123}, nil
}

