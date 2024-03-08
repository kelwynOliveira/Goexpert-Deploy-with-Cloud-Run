run:
	go run ./cmd/main.go

test:
	go test -v ./internal/usecases/

dev:
	docker-compose up -d app_dev

.PHONY: run test dev up