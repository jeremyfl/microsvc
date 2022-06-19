package service

import "gitlab.com/jeremylo/microsvc/grpc/model/stock"

type StockServiceClientMock interface {
	stock.StockServiceClient
}
