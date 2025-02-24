include .envrc
MIGRATIONS_PATH = ./db/migration
test:
	@go test -v ./...
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

migrate-down: 
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

sqlc:
	@sqlc generate

.PHONY: migration migrate-up migrate-down sqlc test