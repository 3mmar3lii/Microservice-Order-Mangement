package main

import (
	"context"
	"errors"
	"fmt"
	"log"

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
		Items:       make([]StockOrderItem, 0),
	}

	if len(items.Items) == 0 {
		return emptyStockResponse, errors.New("order has no items to check")
	}

	productIDs := make([]string, 0, len(items.Items))
	requestedQuantityByID := make(map[string]int64, len(items.Items))
	for _, item := range items.Items {
		if item.Quantity <= 0 {
			return emptyStockResponse, fmt.Errorf("invalid quantity for product %s", item.ID)
		}
		productIDs = append(productIDs, item.ID)
		requestedQuantityByID[item.ID] += int64(item.Quantity)
	}

	query := `SELECT id, name, quantity, unit_price_minor, currency_code
	FROM public.products
	WHERE id = ANY($1);`

	rows, err := s.db.Query(ctx, query, productIDs)
	if err != nil {
		return emptyStockResponse, fmt.Errorf("query products: %w", err)
	}
	defer rows.Close()

	productsByID := make(map[string]StockOrderItem, len(productIDs))
	for rows.Next() {
		var product StockOrderItem
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Quantity,
			&product.UnitPriceMinor,
			&product.CurrencyCode,
		); err != nil {
			return emptyStockResponse, fmt.Errorf("scan product row: %w", err)
		}
		productsByID[product.ID] = product
	}

	if err := rows.Err(); err != nil {
		return emptyStockResponse, fmt.Errorf("iterate product rows: %w", err)
	}

	responseItems := make([]StockOrderItem, 0, len(items.Items))
	for _, item := range items.Items {
		product, exists := productsByID[item.ID]
		if !exists || int64(product.Quantity) < requestedQuantityByID[item.ID] {
			return emptyStockResponse, nil
		}
		product.Quantity = item.Quantity
		responseItems = append(responseItems, product)
	}

	return &StockOrderResponse{
		IsAvailable: true,
		Items:       responseItems,
	}, nil
}
