MIGRATIONS_PATH = ./db/migration
DB_ADDR = postgres://admin:adminpassword@localhost:5432/go_bank?sslmode=disable
test:
	@go test -v -cover ./...
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

migrate-down: 
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

sqlc:
	@sqlc generate

build:
	@go build -o bin/go_bank

server: build
	@./bin/go_bank

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/iangechuki/go_bank/db/sqlc Store
docker-up:
	sudo docker compose up --build
docker-stop:
	sudo docker stop $(sudo docker ps -q)
.PHONY: migration migrate-up migrate-down sqlc test build server mock docker-up docker-stop