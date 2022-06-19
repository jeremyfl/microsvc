package event

import (
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
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
			"order.created",
			stockHandler.OrderCreated,
			"order.created.dlq",
		},
	}
}

