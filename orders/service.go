package main

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/ammar-alaa/common/api"
)

type service struct {
	store       OrderStore
	stockClient pb.StockServiceClient
}

func NewService(store OrderStore, stockClient pb.StockServiceClient) *service {
	return &service{store: store,
		stockClient: stockClient}
}

func (s *service) CreateOrder(ctx context.Context,
	OrderRequest CreateOrderRequest) (Order, error) {

	var stockItems []*pb.ItemsWithQuantity
	for _, item := range OrderRequest.Items {
		stockItems = append(stockItems, &pb.ItemsWithQuantity{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}

	res, err := s.stockClient.CheckOrderInStock(ctx, &pb.CheckIfItemIsInStockRequest{
		Items: stockItems,
	})
	if err != nil {
		return Order{}, err
	}

	if res == nil {
		return Order{}, errors.New("stock service returned an empty response")
	}

	if !res.IsAvailable {
		return Order{}, errors.New("one or more items are unavailable")
	}

	if len(res.Items) != len(OrderRequest.Items) {
		return Order{}, errors.New("stock service returned an unexpected number of items")
	}

	OrderRequest.Items = make([]Item, 0, len(res.Items))
	for _, item := range res.Items {
		if item == nil {
			return Order{}, errors.New("stock service returned an empty item")
		}

		OrderRequest.Items = append(OrderRequest.Items, Item{
			ID:             item.ID,
			Quantity:       item.Quantity,
			Name:           item.Name,
			UnitPriceMinor: item.UnitPriceMinor,
			CurrencyCode:   item.CurrencyCode,
		})
	}

	order, err := s.store.Create(ctx, OrderRequest)
	if err != nil {
		return Order{}, fmt.Errorf("create order: %w", err)
	}

	return order, nil
}
