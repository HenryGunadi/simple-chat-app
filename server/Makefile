build:
	@go build -o bin/chat cmd/main.go

dev: 
	@go run cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/chat

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down