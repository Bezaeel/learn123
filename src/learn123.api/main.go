package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	infra "learn123.infrastructure"
	"learn123.infrastructure/database"
	"learn123.infrastructure/rmq"

	"github.com/gofiber/fiber/v2"

	"learn123.api/modules/course"
	ext "learn123.core/extensions"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load configuration
	cfg, err := ext.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
		os.Exit(1)
	}

	var port = cfg.ServerPort
	infra.AddInfrastucture(&cfg)

	apiHttpServer := NewAPIServer()
	app := apiHttpServer.App()

	// register features
	publisher := rmq.NewEventPublisher(rmq.Channel)
	course.RegisterModule(app, database.DbContext, publisher)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON("Welcome to Learn123 API")
	})

	logger.Info(fmt.Sprintf("starting app on port %v", port))
	logger.Error(app.Listen(fmt.Sprintf(":%v", port)).Error())
}
