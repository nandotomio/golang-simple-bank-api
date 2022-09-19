include .env

# --- Directories and files ---
CURRENT_DIR=$(shell pwd)
MIGRATION_DIR="$(CURRENT_DIR)/db/migrations"
DOCKER_COMPOSE_DEV="docker-compose.dev.yml"

# --- Container Images ---
POSTGRES_IMG="postgres:12.12-alpine3.16"
MIGRATE_IMG="migrate/migrate:v4.15.2"

# --- Default values ---
MIGRATION_NAME="schema"

db-up:
	docker compose -f $(DOCKER_COMPOSE_DEV) up -d postgres

db-down:
	docker compose -f $(DOCKER_COMPOSE_DEV) down postgres

db-drop:
	docker volume rm $$(docker volume ls -qf name=simple_bank_data)

migrate-create:
	docker run --rm -v $(MIGRATION_DIR):/migrations $(MIGRATE_IMG) create -dir /migrations -ext sql -seq "$(MIGRATION_NAME)"

migrate-up:
	docker run --rm -v $(MIGRATION_DIR):/migrations --network host $(MIGRATE_IMG) -path /migrations -database "$(POSTGRES_URL)" -verbose up

migrate-up-1:
	docker run --rm -v $(MIGRATION_DIR):/migrations --network host $(MIGRATE_IMG) -path /migrations -database "$(POSTGRES_URL)" -verbose up 1

migrate-down:
	docker run --rm -v $(MIGRATION_DIR):/migrations --network host $(MIGRATE_IMG) -path /migrations -database "$(POSTGRES_URL)" -verbose down -all

migrate-down-1:
	docker run --rm -v $(MIGRATION_DIR):/migrations --network host $(MIGRATE_IMG) -path /migrations -database "$(POSTGRES_URL)" -verbose down 1

.PHONY: db-up db-down db-drop migrate-create migrate-up migrate-up-1 migrate-down migrate-down-1