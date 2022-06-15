package cmd

import (
	"context"
	"gitlab.com/jeremylo/microsvc/lib"
	"gitlab.com/jeremylo/microsvc/ordersvc/handler/event"
	"log"
)

type consumerTopic struct {
	topic   string
	handler func(message []byte) error
}

func Listen() {
	ctx := context.Background()

	tp := lib.InitTracer("ordersvc-consumer")
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	mb := initMessageBroker(nil, nil)
	db := initDatabase()
	entities := InitEntities(db, mb)

	stockHandler := &event.Handler{
		Services: entities,
	}

	consumers := []consumerTopic{
		{
			"stock.exceeded-amount",
			stockHandler.StockExceededConsumer,
		},
		{
			"product.deleted",
			stockHandler.ProductDeleted,
		},
	}

	for _, consumer := range consumers {
		r := lib.InitMessageReader("stock.exceeded-amount", "ordersvc-consumer")

		for {
			message, err := r.ReadMessage(ctx)
			if err != nil {
				break
			}

			if err = consumer.handler(message.Value); err != nil {
				break
			}

			if err = r.CommitMessages(ctx, message); err != nil {
				log.Println("error when committing", err.Error())
			}
		}

		if err := r.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}

}
