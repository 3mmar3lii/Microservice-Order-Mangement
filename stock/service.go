package main

import (
	"context"
	"log"
)

type Service struct {
	store StockStore
}

func NewStockService(store StockStore) *Service {

	return &Service{}
}

func (s *Service) CheckOrderInStock(ctx context.Context, orderInStock StockOrderRequest) (StockOrderResponse, error) {

	items := []Item{
		{ID: "1", Quantity: 3},
	}
	log.Print("Checking in the stock .....")
	return StockOrderResponse{
		IsAvailable: true,
		Items:       items,
	}, nil
}
