package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/begenov/api-gateway/internal/config"
)

func Open(ctx context.Context, cfg config.PostgresConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
