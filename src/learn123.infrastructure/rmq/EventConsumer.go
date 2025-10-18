package rmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/streadway/amqp"
)

// EventConsumer is a generic consumer interface
type EventConsumer[T any] interface {
	ProcessMessage(ctx context.Context, message T) error
}

// BaseEventConsumer provides base implementation for consuming events
type BaseEventConsumer[T any] struct {
	QueueName    string
	ExchangeName string
	RoutingKey   string
	Logger       *slog.Logger
}

// NewBaseEventConsumer creates a new instance of BaseEventConsumer
func NewBaseEventConsumer[T any](
	queueName string,
	exchangeName string,
	routingKey string,
	logger *slog.Logger,
) *BaseEventConsumer[T] {
	return &BaseEventConsumer[T]{
		QueueName:    queueName,
		ExchangeName: exchangeName,
		RoutingKey:   routingKey,
		Logger:       logger,
	}
}

// Start begins consuming messages from the queue
func (c *BaseEventConsumer[T]) Start(ctx context.Context, ch *amqp.Channel, processor EventConsumer[T]) error {
	// Declare the topic exchange
	err := ch.ExchangeDeclare(
		c.ExchangeName, // exchange name
		"topic",        // exchange type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %v", err)
	}

	queue, err := ch.QueueDeclare(
		c.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	err = ch.QueueBind(
		queue.Name,     // queue name
		c.RoutingKey,   // routing key
		c.ExchangeName, // exchange
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %v", err)
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		return fmt.Errorf("failed to start consuming: %v", err)
	}

	go func() {
		for msg := range msgs {
			var payload T
			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				c.Logger.Error("Failed to unmarshal message", "error", err)
				msg.Nack(false, true)
				continue
			}

			if err := processor.ProcessMessage(ctx, payload); err != nil {
				c.Logger.Error("Failed to process message", "error", err)
				msg.Nack(false, true)
				continue
			}

			msg.Ack(false)
		}
	}()

	return nil
}
