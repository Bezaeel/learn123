ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build:
	@go build ./src/learn123.infrastructure/... && \
	go build ./src/learn123.datamigrator/... && \
	go build ./src/learn123.api/... && \
	go build ./src/learn123.api.client/... && \
	go build ./src/learn123.events/...
	

run:
	@go run src/learn123.datamigrator/*.go && \
	go run src/learn123.api/*.go

add-migration:
	goose -dir ./src/learn123.infrastructure/database/migrations create $(NAME) sql

migrate-up:
	goose -dir ./src/learn123.infrastructure/database/migrations postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" up

migrate-down:
	goose -dir ./src/learn123.infrastructure/database/migrations postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" down