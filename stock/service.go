package main

import (
	"context"
	"log"
)

type Service struct {
	store StockStore
}

func NewStockService(store StockStore) *Service {

	return &Service{store}
}

func (s *Service) CheckOrderInStock(ctx context.Context, orderInStock StockOrderRequest) (*StockOrderResponse, error) {

	log.Print("Checking in the stock .....")

	stockRes, err := s.store.CheckOrderInStockStore(ctx, orderInStock)
	if err != nil {
		return &StockOrderResponse{}, err
	}

	return &StockOrderResponse{
		IsAvailable: stockRes.IsAvailable,
		Items:       stockRes.Items,
	}, nil
}
