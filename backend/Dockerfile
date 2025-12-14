# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Установим git и bash для go modules
RUN apk add --no-cache git bash

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь код
COPY . .

# Сборка бинарника
RUN go build -o /leasing-backend cmd/server/main.go

# Final stage
FROM alpine:3.18

WORKDIR /app

# Копируем бинарник
COPY --from=builder /leasing-backend /leasing-backend

# Создаём папку uploads для временных файлов
RUN mkdir -p ./uploads

EXPOSE 8080

CMD ["/leasing-backend"]
