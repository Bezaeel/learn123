package learn123infrastructure

import (
	_ "github.com/joho/godotenv/autoload"
	ext "learn123.api/common/extensions"
	"learn123.infrastructure/database"
)

var (
	dbUser = ext.EnvString("DB_USER", "")
	dbPass = ext.EnvString("DB_PASS", "")
	dbHost = ext.EnvString("DB_HOST", "")
	dbPort = ext.EnvString("DB_PORT", "")
	dbName = ext.EnvString("DB_NAME", "")
)

func AddInfrastucture() {

	// wire infra dependencies
	database.ConnectToDB(dbUser, dbPass, dbHost, dbPort, dbName)
}
