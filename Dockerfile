FROM golang:latest AS build

WORKDIR /app
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/main.go"]