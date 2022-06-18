package cmd

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gitlab.com/jeremylo/microsvc/lib"
	"gitlab.com/jeremylo/microsvc/ordersvc/handler/event"
	"sync"
)

func consume(consumers []event.ListRoute) {
	var wg sync.WaitGroup
	for _, consumer := range consumers {
		wg.Add(1)
		go func(consumer event.ListRoute, wg *sync.WaitGroup) {
			defer wg.Done()
			r := lib.InitMessageReader(consumer.Topic, Config["APP_NAME"])

			for {
				ctx := context.Background()

				message, err := r.ReadMessage(ctx)
				if err != nil {
					log.WithError(err).Errorln("error when reading messages")

					break
				}

				if err = consumer.Handler(message.Value); err != nil {
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

	tp := lib.InitTracer(Config["APP_NAME"])
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
