package handler

import (
	"context"
	"fmt"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
	"gitlab.com/jeremylo/microsvc/stocksvc/stock"
)

type RpcHandler struct {
	Services domain.Services
}

func (s *RpcHandler) FindStockByProduct(ctx context.Context, message *stock.Message) (*stock.Message, error) {
	return &stock.Message{Body: 123}, nil
}

func (s *RpcHandler) FindEmptyStock(ctx context.Context, message *stock.Message) (*stock.Message, error) {
	fmt.Println(message)

	return &stock.Message{Body: 123}, nil
}
