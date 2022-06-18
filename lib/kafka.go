package lib

import (
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

func InitMessageReader(topic, groupId string) *kafka.Reader {
	host := os.Getenv("KAFKA_HOST")
	if host == "" {
		host = "localhost:9092"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{host},
		Topic:     topic,
		Partition: 0,
		GroupID:   groupId,
	})

	return r
}

func InitMessageWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	}
}
