LOCAL_BIN:=$(CURDIR)/bin
BACKEND_DIR := $(CURDIR)/backend
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.61.0

MIGRATION_FOLDER=$(CURDIR)/data/sql/migrations

ifeq (,$(wildcard .env))
    # Если файл .env отсутствует, устанавливаем параметры по умолчанию
    POSTGRES_USER := 'pg-user'
    POSTGRES_PASSWORD := 'pg-pass'
    POSTGRES_DB := 'larek-dev'
    POSTGRES_HOST := 'localhost'
    POSTGRES_PORT := 5432
else
    # Иначе, подключаем переменные из файла .env
    include .env
    export
endif
POSTGRES_SETUP_TEST := user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=${POSTGRES_HOST} port=${POSTGRES_PORT} sslmode=disable


.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: migration-down
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down


.PHONY: gen-sql
gen-sql:
	sqlc generate