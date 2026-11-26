package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresOrderStore struct {
	pgxpool *pgxpool.Pool
}

func NewPostgresOrderStore(ctx context.Context, pool *pgxpool.Pool) *PostgresOrderStore {
	return &PostgresOrderStore{pool}
}

func (s *PostgresOrderStore) Create(ctx context.Context,OrderRequest CreateOrderRequest ) error {
	// Write query to save order to db 
	// check on the order request data + validate

	return nil
}


