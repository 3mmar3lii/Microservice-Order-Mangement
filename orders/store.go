package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresOrderStore struct {
	pgxpool *pgxpool.Pool
}

func NewPostgresOrderStore(ctx context.Context, pool *pgxpool.Pool) *PostgresOrderStore {
	return &PostgresOrderStore{pool}
}

func (s *PostgresOrderStore) Create(ctx context.Context, request CreateOrderRequest) (Order, error) {
	if strings.TrimSpace(request.CustomerID) == "" {
		return Order{}, errors.New("customer ID is required")
	}

	if len(request.Items) == 0 {
		return Order{}, errors.New("order must contain at least one item")
	}

	for i, item := range request.Items {
		if strings.TrimSpace(item.ID) == "" {
			return Order{}, fmt.Errorf("item %d has an empty ID", i)
		}
		if item.Quantity <= 0 {
			return Order{}, fmt.Errorf("item %q must have a positive quantity", item.ID)
		}
	}

	tx, err := s.pgxpool.Begin(ctx)
	if err != nil {
		return Order{}, fmt.Errorf("begin order transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var order Order
	err = tx.QueryRow(ctx, `
		INSERT INTO public.orders (customerid, status)
		VALUES ($1, 'pending')
		RETURNING id, customerid, status
	`, request.CustomerID).Scan(&order.ID, &order.CustomerID, &order.Status)
	if err != nil {
		return Order{}, fmt.Errorf("insert order: %w", err)
	}

	for _, item := range request.Items {
		_, err = tx.Exec(ctx, `
			INSERT INTO public.order_items (
				order_id, item_id, name, quantity, unit_price_minor, currency_code
			)
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (order_id, item_id)
			DO UPDATE SET
				name = EXCLUDED.name,
				quantity = public.order_items.quantity + EXCLUDED.quantity,
				unit_price_minor = EXCLUDED.unit_price_minor,
				currency_code = EXCLUDED.currency_code
		`, order.ID, item.ID, item.Name, item.Quantity, item.UnitPriceMinor, item.CurrencyCode)
		if err != nil {
			return Order{}, fmt.Errorf("insert order item %q: %w", item.ID, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return Order{}, fmt.Errorf("commit order transaction: %w", err)
	}

	order.Items = append([]Item(nil), request.Items...)
	return order, nil
}
