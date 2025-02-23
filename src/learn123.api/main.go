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
	"github.com/joho/godotenv"

	ext "learn123.api/common/extensions"
	"learn123.api/modules/course"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load .env file
	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	loadErr := godotenv.Load(curDir + "/.env")
	if loadErr != nil {
		log.Fatalln("can't load env file from current directory: " + curDir)
	}

	var port = ext.EnvString("PORT", ":8000")
	infra.AddInfrastucture()

	apiHttpServer := NewAPIServer()
	app := apiHttpServer.App()

	// register features
	publisher := rmq.NewEventPublisher(rmq.Channel)
	course.RegisterModule(app, database.DbContext, publisher)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON("Welcome to Learn123 API")
	})

	logger.Info(fmt.Sprintf("starting app on port %v", port))
	logger.Error(app.Listen(port).Error())
}

func GetAllEnvKeys() ([]string, error) {
	envMap, err := godotenv.Read()
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(envMap))
	for key := range envMap {
		keys = append(keys, key)
	}
	return keys, nil
}
