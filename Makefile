LOCAL_BIN:=$(CURDIR)/bin
BACKEND_DIR := $(CURDIR)/backend # TODO: переименовать в нужный сервис
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.61.0

MIGRATION_FOLDER=data/sql/migrations

ifeq (,$(wildcard .env))
    # Если файл .env отсутствует, устанавливаем параметры по умолчанию
    POSTGRES_USER := cisco
    POSTGRES_PASSWORD := cisco
    POSTGRES_DB := inno-dev
    POSTGRES_HOST := 10.0.1.80
    POSTGRES_PORT := 5432
else
    # Иначе, подключаем переменные из файла .env
    include .env
    export
endif
POSTGRES_SETUP_TEST := user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=${POSTGRES_HOST} port=${POSTGRES_PORT} sslmode=disable
# PG_DSN:=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: sql
sql:
	sqlc generate

.PHONY: lint
lint: # TODO: указать путь до golangci-lint через переменную
	@echo "Starting linter"
	@for dir in $(shell find . -type f -name go.mod -exec dirname {} \;); do \
		echo "Running linter in $$dir"; \
		cd $$dir && golangci-lint run --config $(CURDIR)/.golangci.yml && cd ..; \
	done

.PHONY: run
run:
	@echo "Staring app"
	cd $(BACKEND_DIR) && air
