package handlers

import (
	"context"
	"log/slog"

	"learn123.infrastructure/rmq"

	"github.com/streadway/amqp"
	event "learn123.events"
)

type OrderCreatedConsumer struct {
	base   *rmq.BaseEventConsumer[event.OrderCreated]
	logger *slog.Logger
}

func NewOrderCreatedConsumer(logger *slog.Logger) *OrderCreatedConsumer {
	return &OrderCreatedConsumer{
		base: rmq.NewBaseEventConsumer[event.OrderCreated](
			"learn123.OrderCreated",
			"learn123",
			"learn123.OrderCreated",
			logger,
		),
		logger: logger,
	}
}

func (c *OrderCreatedConsumer) ProcessMessage(ctx context.Context, message event.OrderCreated) error {
	c.logger.Info("Processing order created event",
		"orderID", message.Id,
		"customerID", message.CreatedBy,
	)
	// Add your order processing logic here
	return nil
}

func (c *OrderCreatedConsumer) Start(ctx context.Context, ch *amqp.Channel) error {
	return c.base.Start(ctx, ch, c)
}
