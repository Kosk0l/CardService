package main

import (
	"CardService/internal/grpchandler"
	"CardService/internal/services"
	"CardService/internal/storage"
	pb "CardService/proto/grpcProto"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	// 1. объект подключения к БД (репозиторий)
    pool, err := storage.NewPostgres(context.Background(), "postgres")
    if err != nil {
        log.Fatalf("DB connect error: %v", err)
    }

	// Объект storage
    repo := pool  // repo реализует CardRepository

	// 2. объект Бизнес-логики
    service := services.NewService(repo)

    // 3. объект gRPC хендлера, зависящий от сервиса
    handler := grpchandler.NewServer(service)


	// Настройка слушателя
	lis, err := net.Listen("tcp", ":44045")
	if err != nil {
		log.Fatalf("Error listen server: %v", err)
	}

	grpcServer := grpc.NewServer()
    pb.RegisterCardServiceServer(grpcServer, handler)

    log.Println("gRPC server started on :44045")

    // 5. Старт сервера
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}