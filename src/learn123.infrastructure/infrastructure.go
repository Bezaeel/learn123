package learn123infrastructure

import (
	"log"
	"os"

	ext "learn123.core/extensions"
	"learn123.infrastructure/database"
	"learn123.infrastructure/rmq"
)

func AddInfrastucture(config *ext.Config) {
	var cfg ext.Config
	if config == nil {
		var err error
		cfg, err = ext.LoadConfig()
		if err != nil {
			log.Fatalf("cannot load config: %v", err)
			os.Exit(1)
		}
	} else {
		cfg = *config
	}

	// Wire infra dependencies
	database.ConnectToDB(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	rmq.ConnectAmqp()
}
