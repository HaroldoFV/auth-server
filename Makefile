include .env
export

DB_URL := $(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_HOST)/$(DB_NAME)?sslmode=disable

createmigration:
	migrate create -ext=sql -dir=internal/adapters/outgoing/migrations -seq init

migrate:
	migrate -path=internal/adapters/outgoing/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path=internal/adapters/outgoing/migrations -database "$(DB_URL)" -verbose down

.PHONY: migrate migratedown createmigration