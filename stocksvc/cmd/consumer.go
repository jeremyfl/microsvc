package cmd

import (
	"context"
	"gitlab.com/jeremylo/microsvc/lib"
	"gitlab.com/jeremylo/microsvc/stocksvc/handler/event"
	"log"
	"sync"
)

func consume(consumers []event.ListRoute) {
	var wg sync.WaitGroup
	for _, consumer := range consumers {
		wg.Add(1)
		go func(consumer event.ListRoute, wg *sync.WaitGroup) {
			defer wg.Done()
			r := lib.InitMessageReader(consumer.Topic, lib.Config["APP_NAME"])

			for {
				ctx := context.Background()

				message, err := r.ReadMessage(ctx)
				if err != nil {
					log.Println(err.Error())

					break
				}

				if err = consumer.Handler(message.Value); err != nil {
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

	tp := lib.InitTracer(lib.Config["APP_NAME"])
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	w := lib.InitMessageWriter()
	mb := initMessageBroker(w, nil)
	db := initDatabase()
	entities := InitEntities(db, mb)
	route := event.Route(entities)

	consume(route)
}
