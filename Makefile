.PHONY: build test vet tidy migrate seed

build:
	go build ./...

test:
	go test ./...

vet:
	go vet ./...

tidy:
	go mod tidy

migrate:
	go run ./cmd/migrate

seed:
	go run ./cmd/seed
