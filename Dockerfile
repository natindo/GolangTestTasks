FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tasks-api ./cmd/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/tasks-api /app/tasks-api

EXPOSE 8080

ENTRYPOINT ["./tasks-api"]
