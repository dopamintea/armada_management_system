FROM golang:1.24.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /api ./cmd/api/main.go
RUN go build -o /mqtt_listener ./cmd/mqtt_listener/main.go
RUN go build -o /worker ./cmd/worker/main.go
RUN go build -o /mock_publisher ./cmd/mock_publisher/main.go

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /api .
COPY --from=builder /mqtt_listener .
COPY --from=builder /worker .
COPY --from=builder /mock_publisher .

COPY .env-example .env

CMD ["./api"]
