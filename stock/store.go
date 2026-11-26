package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStockStore struct {
	db *pgxpool.Pool
}

func NewPostgresStore(db *pgxpool.Pool) *PostgresStockStore {
	return &PostgresStockStore{db}
}

func (s *PostgresStockStore) CheckOrderInStockStore(ctx context.Context, items StockOrderRequest) (*StockOrderResponse, error) {

	log.Printf("Checking in the stock store ..... %v", items)

	emptyStockResponse := &StockOrderResponse{
		IsAvailable: false,
		Items:       make([]Item, 0),
	}

	if len(items.Items) == 0 {
		return emptyStockResponse, errors.New("order has no items to check")
	}

	query := `SELECT quantity FROM products WHERE id = $1;`

	for _, item := range items.Items {
		if item.Quantity <= 0 {
			return emptyStockResponse, fmt.Errorf("invalid quantity for product %s", item.ID)
		}

		var availableQty int
		err := s.db.QueryRow(ctx, query, item.ID).Scan(&availableQty)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return emptyStockResponse, nil
			}
			return emptyStockResponse, fmt.Errorf("query product %s: %w", item.ID, err)
		}

		if availableQty < int(item.Quantity) {
			return emptyStockResponse, nil
		}
	}

	return &StockOrderResponse{
		IsAvailable: true,
		Items:       items.Items,
	}, nil
}
