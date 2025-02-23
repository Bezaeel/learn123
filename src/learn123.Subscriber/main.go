package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	event "learn123.events"
	"learn123.infrastructure/rmq"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	rmq.ConnectAmqp()
	logger.Info("Subscriber started")
	queue, _ := rmq.Channel.QueueDeclare(
		"learn123.CourseCreated", // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)

	_ = rmq.Channel.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"learn123", // exchange
		false,
		nil,
	)

	msgs, err := rmq.Channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        //args
	)
	if err != nil {
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			o := &event.CourseCreated{}
			if err := json.Unmarshal(msg.Body, o); err != nil {
				msg.Nack(false, true)
				logger.Error("Failed to unmarshal message: %s\n", err)
				continue
			}
			fmt.Println(fmt.Sprintf("Received a message: %s", o.Name))
		} 
	}()

	logger.Info("Waiting for messages...")
	<-forever

}
