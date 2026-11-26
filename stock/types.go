package main

import "context"

type StockOrderRequestItem struct {
	ID       string
	Quantity int32
}

type StockOrderRequest struct {
	Items []StockOrderRequestItem
}

type StockOrderItem struct {
	ID             string
	Quantity       int32
	UnitPriceMinor int64
	CurrencyCode   string
}

type StockOrderResponse struct {
	IsAvailable bool
	Items       []StockOrderItem
}

// Buss logic
type StockService interface {
	CheckOrderInStock(context.Context, StockOrderRequest) (*StockOrderResponse, error)
}

// DB
type StockStore interface {
	CheckOrderInStockStore(context.Context, StockOrderRequest) (*StockOrderResponse, error)
}
