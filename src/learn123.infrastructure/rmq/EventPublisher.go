package rmq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
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

func (p *EventPublisher) Publish(eventType string, event interface{}) error {
	// Use the provided eventType as routing key
	routingKey := fmt.Sprintf("learn123.%s", eventType)

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}
	fmt.Printf("Publishing event %v, %v", routingKey, string(body))

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
