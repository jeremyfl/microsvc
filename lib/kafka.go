package lib

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

var kafkaHost string

func init() {
	kafkaHost = os.Getenv("KAFKA_HOST")

	if kafkaHost == "" {
		kafkaHost = "localhost"
	}

	kafkaHost = fmt.Sprintf("%s:9092", kafkaHost)
}

func InitMessageReader(topic, groupId string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaHost},
		Topic:     topic,
		Partition: 0,
		GroupID:   groupId,
	})

	return r
}

func InitMessageWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(kafkaHost),
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	}
}
