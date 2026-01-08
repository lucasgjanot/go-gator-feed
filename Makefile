.PHONY: service-up service-down migrations-up wait-db database-up 

ENV_FILE := .env.development
DB_CONTAINER := postgres-dev


# carrega o .env.development para o make
ifneq (,$(wildcard $(ENV_FILE)))
	include $(ENV_FILE)
	export
endif

DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

service-up: database-up wait-db migrations-up

service-down:
	docker compose -f ./infra/compose.yaml --env-file $(ENV_FILE) down

database-up:
	docker compose -f ./infra/compose.yaml --env-file $(ENV_FILE) up -d

wait-db:
	@echo -n "⏳ Waiting for database"
	@until docker exec \
		$(DB_CONTAINER) \
		pg_isready \
			-U $(POSTGRES_USER) \
			-d $(POSTGRES_DB) > /dev/null 2>&1; do \
		printf "."; \
		sleep 1; \
	done
	@echo ""
	@echo "✅ Database is ready"


migrations-up:
	@echo "🚀 Running migrations..."
	goose -dir sql/schema postgres "$$DATABASE_URL" up

