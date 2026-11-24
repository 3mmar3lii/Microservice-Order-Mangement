package main

import "context"

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(ctx context.Context,
	OrderRequest CreateOrderRequest) (Order, error) {

	return Order{
		ID: "20",
	},nil 
}
