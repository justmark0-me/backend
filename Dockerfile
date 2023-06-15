FROM golang:1.20-alpine

WORKDIR /app

COPY . .

CMD go run src/main.go
