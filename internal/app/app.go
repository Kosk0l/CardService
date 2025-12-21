package app

import (
	"CardService/config"
	"CardService/internal/grpchandler"
	"CardService/internal/server"
	"CardService/internal/services"
	"CardService/internal/storage"
	"context"
	"fmt"
)

type App struct {
	grpc *server.GRPCServer
}

// Связка компонентов // Конструктор структуры App
func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Pass,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	// Объект storage
	pool, err := storage.NewPostgres(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// Объект бизнес-логики
	repo := pool
	svc := services.NewService(repo)

	// Объект gRPC handler
	handler := grpchandler.NewServer(svc)

	// Объект сервера
	grpcServer := server.NewGrpcServer(cfg.App.Port, handler)

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