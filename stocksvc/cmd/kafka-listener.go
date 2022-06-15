package cmd

import (
	"context"
	"fmt"
	"gitlab.com/jeremylo/microsvc/stocksvc/domain"
	"gitlab.com/jeremylo/microsvc/stocksvc/handler"
	"log"
)



func listener(entities domain.Services) {
	r := initMessageReader()
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
