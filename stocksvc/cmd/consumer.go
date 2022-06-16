package cmd

import (
	"context"
	"gitlab.com/jeremylo/microsvc/lib"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
	"gitlab.com/jeremylo/microsvc/stocksvc/handler/event"
	"log"
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
			"order.created",
			stockHandler.OrderCreated,
			"order.created.dlq",
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
					log.Println(err.Error())

					break
				}

				if err = consumer.handler(message.Value); err != nil {
					log.Println(err.Error())

					break
				}

				if err = r.CommitMessages(ctx, message); err != nil {
					log.Println("error when committing", err.Error())
				}
			}

			if err := r.Close(); err != nil {
				log.Fatal("failed to close reader:", err)
			}

			wg.Done()
		}(consumer, &wg)
	}

	wg.Wait()
}

func Listen() {
	ctx := context.Background()

	tp := lib.InitTracer("stocksvc-consumer")
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
