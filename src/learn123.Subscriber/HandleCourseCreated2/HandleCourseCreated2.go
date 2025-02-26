package HandleCourseCreated2

import (
	"context"
	"log/slog"

	"learn123.infrastructure/rmq"

	"github.com/streadway/amqp"
	event "learn123.events"
)

type CourseCreatedConsumer struct {
	base   *rmq.BaseEventConsumer[event.CourseCreated]
	logger *slog.Logger
}

func NewCourseCreatedConsumer(logger *slog.Logger) *CourseCreatedConsumer {
	return &CourseCreatedConsumer{
		base: rmq.NewBaseEventConsumer[event.CourseCreated](
			"learn123.CourseCreated.subscriber2",
			"learn123",
			"learn123.CourseCreated",
			logger,
		),
		logger: logger,
	}
}

func (c *CourseCreatedConsumer) ProcessMessage(ctx context.Context, message event.CourseCreated) error {
	c.logger.Info("Subscriber 2: Processing course created event", "courseName", message.Name)
	// Add your business logic here
	return nil
}

func (c *CourseCreatedConsumer) Start(ctx context.Context, ch *amqp.Channel) error {
	return c.base.Start(ctx, ch, c)
}
