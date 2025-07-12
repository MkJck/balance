package database

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context) (*pgxpool.Pool, error) {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        // Пример: "postgres://user:password@localhost:5432/balance?sslmode=disable"
        dsn = "postgres://postgres:postgres@localhost:5432/balance?sslmode=disable"
    }
    pool, err := pgxpool.New(ctx, dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to postgres: %w", err)
    }
    return pool, nil
}