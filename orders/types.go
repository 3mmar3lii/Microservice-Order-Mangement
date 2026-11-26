package main

import (

	"context"
)

type Item struct {
	ID       string
	Quantity int32
}

type CreateOrderRequest struct {
	CustomerID string
	Items      []Item
}

type Order struct {
	ID         string
	CustomerID string
	Status     string
	Items      []Item
}
// Store Interface for Service layer - Bussiness
type OrderService interface {
	CreateOrder(context.Context, CreateOrderRequest) (Order, error)
}

// Store Interface for DB
type OrderStore interface {
	Create(context.Context, CreateOrderRequest) error
}
