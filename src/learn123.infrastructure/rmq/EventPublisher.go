package rmq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	ext "learn123.core/extensions"

)

type EventPublisher struct {
	channel *amqp.Channel
}

func NewEventPublisher(ch *amqp.Channel) *EventPublisher {
	// Declare the topic exchange
	err := ch.ExchangeDeclare(
		"learn123", // exchange name
		"topic",    // exchange type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil
	}

	return &EventPublisher{
		channel: ch,
	}
}

func (p *EventPublisher) Publish(event interface{}) error {
	// Get the type name and use it as routing key
	// Convert type name to routing key format
	// e.g., "CourseCreated" -> "learn123.CourseCreated"
	routingKey := fmt.Sprintf("learn123.%s", ext.GetType(&event))

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.channel.Publish(
		"learn123", // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
