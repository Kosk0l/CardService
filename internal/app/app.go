package app

import (
	"CardService/config"
	"CardService/internal/grpchandler"
	"CardService/internal/server"
	"CardService/internal/services"
	"CardService/internal/storage"
	"context"
)

type App struct {
	grpc *server.GRPCServer
}

// Связка компонентов // Конструктор структуры App
func NewApp(ctx context.Context, cfg config.Config) (*App, error) {
	// Объект storage
	pool, err := storage.NewPostgres(ctx, cfg.DataBaseURL)
	if err != nil {
		return nil, err
	}

	// Объект бизнес-логики
	repo := pool
	services := services.NewService(repo)

	// Объект gRPC handler
	handler := grpchandler.NewServer(services)

	// Объект сервера
	grpcServer := server.NewGrpcServer(cfg.GRPCPort, handler)

	return &App{
		grpc: grpcServer,
	}, nil
}

func (s *App) Run() (error) {
	return s.grpc.Start()
}

func (s *App) Stop() () {
	s.grpc.Stop()
}