package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sirupsen/logrus"
	"learn123.Subscriber2/handlers"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Create a new AMQP subscriber
	amqpURI := "amqp://guest:guest@localhost:5672/"
	subscriber, err := amqp.NewSubscriber(
		amqp.NewDurablePubSubConfig(
			amqpURI,
			func(topic string) string {
				return "learn123." + topic + ".subscriber2"
			},
		),
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		log.Fatalf("Failed to create subscriber: %v", err)
	}

	// Create a new message router
	router, err := message.NewRouter(message.RouterConfig{}, watermill.NewStdLogger(false, false))
	if err != nil {
		log.Fatalf("Failed to create router: %v", err)
	}

	// Create handlers
	courseHandler := handlers.NewCourseCreatedHandler(logger)
	orderHandler := handlers.NewOrderCreatedHandler(logger)

	// Add handlers to router
	router.AddHandler(
		"course_created_handler",
		"CourseCreated",
		subscriber,
		"",
		nil,
		courseHandler.Handle,
	)

	router.AddHandler(
		"order_created_handler",
		"OrderCreated",
		subscriber,
		"",
		nil,
		orderHandler.Handle,
	)

	// Create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Run the router
	if err := router.Run(ctx); err != nil {
		log.Fatalf("Failed to run router: %v", err)
	}
}
