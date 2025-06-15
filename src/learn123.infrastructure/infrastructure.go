package learn123infrastructure

import (
	ext "learn123.core/extensions"
	"learn123.infrastructure/database"
	"learn123.infrastructure/rmq"
)

var (
	dbUser = ext.EnvString("DB_USER", "")
	dbPass = ext.EnvString("DB_PASS", "")
	dbHost = ext.EnvString("DB_HOST", "")
	dbPort = ext.EnvString("DB_PORT", "")
	dbName = ext.EnvString("DB_NAME", "")
)

func AddInfrastucture() {
	// Wire infra dependencies
	database.ConnectToDB(dbUser, dbPass, dbHost, dbPort, dbName)
	rmq.ConnectAmqp()
}
