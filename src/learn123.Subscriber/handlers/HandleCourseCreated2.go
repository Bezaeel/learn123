package handlers

import (
	"context"
	"log/slog"

	"learn123.infrastructure/rmq"

	"github.com/streadway/amqp"
	event "learn123.events"
)

type CourseCreatedConsumer2 struct {
	base   *rmq.BaseEventConsumer[event.CourseCreated]
	logger *slog.Logger
}

func NewCourseCreatedConsumer2(logger *slog.Logger) *CourseCreatedConsumer2 {
	return &CourseCreatedConsumer2{
		base: rmq.NewBaseEventConsumer[event.CourseCreated](
			"learn123.CourseCreated.subscriber2",
			"learn123",
			"learn123.CourseCreated",
			logger,
		),
		logger: logger,
	}
}

func (c *CourseCreatedConsumer2) ProcessMessage(ctx context.Context, message event.CourseCreated) error {
	c.logger.Info("Subscriber: Processing course 2 created event", "courseName", message.Name)
	// Add your business logic here
	return nil
}

func (c *CourseCreatedConsumer2) Start(ctx context.Context, ch *amqp.Channel) error {
	return c.base.Start(ctx, ch, c)
}
