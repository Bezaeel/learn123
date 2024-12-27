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
	@go run src/learn123.datamigrator/main.go && \
	go run src/learn123.api/*.go