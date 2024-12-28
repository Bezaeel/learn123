package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"
	infra "learn123.infrastructure"
	"learn123.infrastructure/database"
)

var migrationPath = "./src/learn123.infrastructure/database/migrations"

func main() {
	infra.AddInfrastucture()
	fmt.Println("Running migration...")

	// Read migration files from the directory
	files, err := os.ReadDir(migrationPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			fmt.Println("Found migration file:", file.Name())
		}
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	sqlDB, err := database.DbContext.DB()
	if err != nil {
		panic(err)
	}

	if err := goose.Up(sqlDB, migrationPath); err != nil {
		panic(err)
	}
}
