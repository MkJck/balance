package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"balance/internal/config"

	"github.com/jackc/pgx/v5"
)

func Connect(cfg config.DatabaseConfig) (*pgx.Conn, error) {
	// Формируем строку подключения
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode)

	// Подключение к базе с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверяем подключение
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("PostgreSQL successfully connected!")
	return db, nil
}
