package stock

import (
	"context"
	"fmt"
	"stocksvc/domain"
)

type Server struct {
	Services domain.Services
}

func (s *Server) FindStockByProduct(ctx context.Context, message *Message) (*Message, error) {
	return &Message{Body: 123}, nil
}

func (s *Server) FindEmptyStock(ctx context.Context, message *Message) (*Message, error) {
	fmt.Println(message)

	return &Message{Body: 123}, nil
}

