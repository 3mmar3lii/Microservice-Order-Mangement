package main

import (
	"context"

	pb "github.com/ammar-alaa/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	pb.UnimplementedStockServiceServer
	service StockService
}

func NewGRPCHandler(grpcServer *grpc.Server, svc StockService) {
	handler := &grpcHandler{
		service: svc,
	}
	pb.RegisterStockServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CheckOrderInStock(ctx context.Context,
	orderRequest *pb.CheckIfItemIsInStockRequest) (*pb.CheckIfItemIsInStockResponse, error) {
	if orderRequest == nil {
		return nil, status.Error(codes.InvalidArgument, "order request is required")
	}

	items := make([]StockOrderRequestItem, 0, len(orderRequest.Items))
	for _, item := range orderRequest.Items {
		if item == nil {
			return nil, status.Error(codes.InvalidArgument, "order request contains a nil item")
		}
		items = append(items, StockOrderRequestItem{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}

	stockResponse, err := h.service.CheckOrderInStock(ctx, StockOrderRequest{Items: items})
	if err != nil {
		return nil, err
	}

	responseItems := make([]*pb.Items, 0, len(stockResponse.Items))
	for _, item := range stockResponse.Items {
		responseItems = append(responseItems, &pb.Items{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}

	return &pb.CheckIfItemIsInStockResponse{
		IsAvailable: stockResponse.IsAvailable,
		Items:       responseItems,
	}, nil
}
