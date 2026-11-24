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

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(c context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {

	log.Printf("New Order recived! Order %v", p)
	o := &pb.Order{
		ID: "20",
	}
	return o, nil
}
