run:
	go run ./cmd/main.go

test:
	go test -v ./internal/usecases/

docker:
	docker-compose up --build -d

.PHONY: run test dev up