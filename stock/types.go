package main

import "context"

type Item struct {
	ID       string
	Quantity int32
}

type StockOrderRequest struct {
	Items []Item
}

type StockOrderResponse struct {
	IsAvailable bool
	Items       []Item
}

// Buss logic
type StockService interface {
	CheckOrderInStock(context.Context, StockOrderRequest) (*StockOrderResponse, error)
}

// DB
type StockStore interface {
	CheckOrderInStockStore(context.Context, StockOrderRequest) (*StockOrderResponse, error)
}
