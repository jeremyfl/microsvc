package grpc

import (
	"context"
	"fmt"
	"gitlab.com/jeremylo/microsvc/grpc/model/stock"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
)

type RpcHandler struct {
	Services domain.Services
}

func (s *RpcHandler) FindStockByProduct(ctx context.Context, message *stock.Message) (*stock.Message, error) {
	return &stock.Message{Body: message.Body}, nil
}

func (s *RpcHandler) FindEmptyStock(ctx context.Context, message *stock.Message) (*stock.Message, error) {
	fmt.Println(message)

	return &stock.Message{Body: 123}, nil
}
