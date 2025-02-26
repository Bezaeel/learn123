package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"learn123.infrastructure/rmq"
	"learn123.subscriber/HandleCourseCreated"
	"learn123.subscriber/HandleCourseCreated2"
	"learn123.subscriber/HandleOrderCreated"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Setup RabbitMQ connection
	rmq.ConnectAmqp()

	// Create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initialize and start consumers
	courseCreatedConsumer := HandleCourseCreated.NewCourseCreatedConsumer(logger)
	if err := courseCreatedConsumer.Start(ctx, rmq.Channel); err != nil {
		logger.Error("Failed to start course created consumer", "error", err)
		os.Exit(1)
	}

	orderCreatedConsumer := HandleOrderCreated.NewOrderCreatedConsumer(logger)
	if err := orderCreatedConsumer.Start(ctx, rmq.Channel); err != nil {
		logger.Error("Failed to start order created consumer", "error", err)
		os.Exit(1)
	}

	courseCreatedConsumer2 := HandleCourseCreated2.NewCourseCreatedConsumer(logger)
	if err := courseCreatedConsumer2.Start(ctx, rmq.Channel); err != nil {
		logger.Error("Failed to start course created2 consumer", "error", err)
		os.Exit(1)
	}

	logger.Info("Service started. Press CTRL+C to exit")

	// Wait for interrupt signal
	<-ctx.Done()
	logger.Info("Shutting down gracefully")
}
