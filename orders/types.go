package main

import "context"


// Store Interface for Service layer - Bussiness
type OrderService interface {
	CreateOrder(context.Context) error
}

// Store Interface for DB
type OrderStore interface {
	Create (context.Context) error
}

