package cmd

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"stocksvc/domain"
	"stocksvc/handler"
)

func newReader() *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "order.created",
		Partition: 0,
		GroupID:   "stocksvc-listener",
	})

	return r
}

func listener(entities domain.Services) {
	r := newReader()
	h := handler.OrderCreatedHandler{Services: entities}

	fmt.Println("Listening to order message")

	for {
		ctx := context.Background()

		message, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println(err.Error())
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
