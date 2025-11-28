package storage

import (
	"context"
	"fmt"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
	"CardService/internal/services"
)

type Postgres struct {
	pool *pgxpool.Pool
	services.CardRepository // Имплемент интерфейса из высокого уровня
}

// Создание нового соединения к БД
func NewPostgres(ctx context.Context, dsn string) (*Postgres, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed parse Config Postgres pgxpool: %v", err)
	}

	// Настройка пула и таймаута
	config.MaxConns = 10 // Максимальное количество открытых соединений
	config.MaxConnLifetime = time.Hour // Максимальное время жизни соединения

	// Создание пула соединений
	newpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("ailed config connection: %v", err)
	}

	return &Postgres{
		pool: newpool,
	}, nil
}

// Закрывает пул соединений
func (p *Postgres) Close() {
	p.pool.Close()
}
