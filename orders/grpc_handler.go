package main

import (
	"context"
	"log"

	pb "github.com/ammar-alaa/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrderService
}

func NewGRPCHandler(grpcServer *grpc.Server, svc OrderService) {
	handler := &grpcHandler{
		service: svc,
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(c context.Context, orderRequest *pb.CreateOrderRequest) (*pb.Order, error) {

	var items []Item
	log.Printf("New Order recived! Order %v", orderRequest)

	// Map the Contract to the normal types we have first
	for _, item := range orderRequest.Items {
		items = append(items, Item{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}
	req := CreateOrderRequest{
		CustomerID: orderRequest.CustomerID,
		Items:      items,
	}
	// Delegate work to the order service real buss logic
	order, err :=h.service.CreateOrder(c, req)
	if err != nil {
		return nil,err
	}
	
	return &pb.Order{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Status:     order.Status,
	}, nil
}
