package main

import (
	"fmt"
	"log/slog"
	"os"

	infra "learn123.infrastructure"
	"learn123.infrastructure/database"

	"github.com/gofiber/fiber/v2"

	ext "learn123.api/common/extensions"
	"learn123.api/modules/course"
)

func main() {
	var port = ext.EnvString("PORT", ":8000")
	infra.AddInfrastucture()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	apiHttpServer := NewAPIServer()
	app := apiHttpServer.App()

	// register features
	course.RegisterModule(app, database.DbContext)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON("Welcome to Learn123 API")
	})

	logger.Info(fmt.Sprintf("starting app on port %v", port))
	logger.Error(app.Listen(port).Error())
}
