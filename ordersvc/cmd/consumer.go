package cmd

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/jeremylo/microsvc/lib"
	"gitlab.com/jeremylo/microsvc/ordersvc/domain"
	"gitlab.com/jeremylo/microsvc/ordersvc/handler/event"
	"sync"
)

type consumerTopic struct {
	topic    string
	handler  func(message []byte) error
	dlqTopic string
}

func consumersRoute(entities domain.Services) []consumerTopic {
	stockHandler := &event.Handler{Services: entities}

	return []consumerTopic{
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

func consume(consumers []consumerTopic) {
	var wg sync.WaitGroup
	for _, consumer := range consumers {
		wg.Add(1)
		go func(consumer consumerTopic, wg *sync.WaitGroup) {
			defer wg.Done()
			r := lib.InitMessageReader(consumer.topic, "ordersvc-consumer")

			for {
				ctx := context.Background()

				message, err := r.ReadMessage(ctx)
				if err != nil {
					log.WithError(err).Errorln("error when reading messages")

					break
				}

				if err = consumer.handler(message.Value); err != nil {
					log.WithError(err).Errorln("error when handling the message from handler")

					break
				}
			}

			if err := r.Close(); err != nil {
				log.WithError(err).Errorln("failed to close reader")
			}

			wg.Done()
		}(consumer, &wg)
	}

	wg.Wait()
}

func Listen() {
	ctx := context.Background()

	tp := lib.InitTracer("ordersvc-consumer")
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	w := lib.InitMessageWriter()
	mb := initMessageBroker(w, nil)
	db := initDatabase()
	entities := InitEntities(db, mb)
	route := consumersRoute(entities)

	consume(route)
}
