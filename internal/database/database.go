package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/dunamismax/go-modern-scaffold/internal/database/sqlc"
)

// DB represents the database connection.
type DB struct {
	*sqlc.Queries
	pool *pgxpool.Pool
}

// New creates a new database connection and returns a DB instance.
func New(dsn string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create database connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{
		Queries: sqlc.New(pool),
		pool:    pool,
	}, nil
}

// Close closes the database connection.
func (db *DB) Close() {
	db.pool.Close()
}
