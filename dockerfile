# STAGE 1: build 
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# 1. Копирует только go.mod / go.sum
COPY go.mod go.sum ./
RUN go mod download

# 2. Копирует весь проект
COPY . .

# 3. Собирает бинарник cardservice
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o cardservice ./cmd/main

# STAGE 2: runtime 
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# 4. Копирует бинарник из builder
COPY --from=builder /app/cardservice /app/cardservice

# 5. Указывает порт
EXPOSE 8080

# 6. Запуск
CMD ["/app/cardservice"]

# docker build -t cardservice .