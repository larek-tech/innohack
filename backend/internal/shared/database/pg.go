package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/larek-tech/innohack/backend/internal/auth/storage"
	"github.com/larek-tech/innohack/backend/pkg/pg"
)

type PostgresConfig struct {
	DSN string `yaml:"dsn"`
}

func InitPostgres(ctx context.Context, connStr string) *storage.PG {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	q := pg.New(pool)

	return storage.NewPG(pool, q)
}
