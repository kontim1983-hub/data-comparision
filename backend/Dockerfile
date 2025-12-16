# Stage 1: Build
# Используем golang 1.25 alpine и даём alias "builder"
FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /leasing-backend cmd/server/main.go

# Stage 2: Production
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /leasing-backend /leasing-backend
RUN chmod +x /leasing-backend
RUN mkdir -p ./uploads

EXPOSE 8080

CMD ["/leasing-backend"]