package event

import (
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
)

type ListRoute struct {
	Topic    string
	Handler  func(message []byte) error
	DlqTopic string
}

func Route(entities domain.Services) []ListRoute {
	stockHandler := &Handler{Services: entities}

	return []ListRoute{
		{
			"stock.exceeded-amount",
			stockHandler.StockExceededConsumer,
			"stock.exceeded-amount.dlq",
		},
		{
			"product.deleted",
			stockHandler.GenericHandler,
			"product.deleted.dlq",
		},
	}
}

