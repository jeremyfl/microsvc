package domain

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type MessageBrokerImpl struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
}

type MessageBroker interface {
	Publish(ctx context.Context, topic string, message interface{}) error
	Subscribe(topic string, handler func(message interface{})) error
}

func (m *MessageBrokerImpl) Publish(ctx context.Context, topic string, message interface{}) error {
	p, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return m.Writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: p,
	})
}

func (m *MessageBrokerImpl) Subscribe(topic string, handler func(message interface{})) error {
	//TODO implement me
	panic("implement me")
}