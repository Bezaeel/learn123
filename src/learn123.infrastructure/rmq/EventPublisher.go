package rmq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)


type EventPublisher struct {
	exchange    string
	routingKey  string
	amqpChannel *amqp.Channel
}

func NewEventPublisher(amqpChannel *amqp.Channel) *EventPublisher {
	return &EventPublisher{
		exchange:    "learn123",
		routingKey:  "",
		amqpChannel: amqpChannel,
	}
}

func (e *EventPublisher) Publish(event interface{}) {
	jsonEvent, _ := json.Marshal(event)
	_ = e.amqpChannel.Publish(
		e.exchange, // exchange
		"",
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonEvent,
		},
	)
	fmt.Println(fmt.Sprintf("Event published: %s", jsonEvent))
}
