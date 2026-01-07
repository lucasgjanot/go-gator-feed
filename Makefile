.PHONY: service-up service-down migrate-up wait-db database-up

ENV_FILE := .env.development
DB_CONTAINER := postgres-dev


# carrega o .env.development para o make
ifneq (,$(wildcard $(ENV_FILE)))
	include $(ENV_FILE)
	export
endif

DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

service-up: database-up wait-db migrate-up

database-up:
	docker compose -f ./infra/compose.yaml --env-file $(ENV_FILE) up -d

wait-db:
	@echo "⏳ Waiting for database..."
	@until docker exec --env-file $(ENV_FILE) \
		$(DB_CONTAINER) \
		pg_isready \
		> /dev/null 2>&1; do \
		sleep 1; \
	done
	@echo "✅ Database is ready"

migrate-up:
	@echo "🚀 Running migrations..."
	goose -dir sql/schema postgres "$$DATABASE_URL" up

service-down:
	docker compose -f ./infra/compose.yaml --env-file $(ENV_FILE) down
