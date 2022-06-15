package cmd

import (
	"context"
	"fmt"
	"gitlab.com/jeremylo/microsvc/lib"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
	"gitlab.com/jeremylo/microsvc/stocksvc/handler"
	"log"
)

func Listen(entities domain.Services) {
	tp := lib.InitTracer("stocksvc-listener")
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	r := lib.InitMessageReader("order.created", "stocksvc-listener")
	h := handler.OrderCreatedHandler{Services: entities}

	fmt.Println("Listening to order message")

	for {
		ctx := context.Background()

		message, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}

		if err = h.Handle(ctx, message.Value); err != nil {
			log.Println("error when processing", err.Error())

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
