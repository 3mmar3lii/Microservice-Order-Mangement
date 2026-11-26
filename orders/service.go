package main

import (
	"context"

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

	if res == nil || len(res.Items) == 0 || res.Items[0] == nil {
		return Order{}, nil
	}

	// Call Order Store to truly store Orde with status pending 
	// in postgress order table
	if res.IsAvailable {
		err :=s.store.Create(ctx,OrderRequest)
		if err != nil {
			// write good error and return it explcicty 
		}
	}

	return Order{ID: res.Items[0].ID}, nil
}
