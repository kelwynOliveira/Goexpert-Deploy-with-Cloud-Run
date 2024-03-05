docker:
	docker-compose up -d

run:
	go run ./cmd/main.go

test:
	go test -v ./internal/usecases/

dev:
	docker-compose up api_dev

.PHONY: run test dev