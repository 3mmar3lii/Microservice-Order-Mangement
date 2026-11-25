package main

import (
	"context"
	"log"
	"net"

	"github.com/ammar-alaa/common"
	"google.golang.org/grpc"
)

var (
	OrderAddrs = common.EnvString("OrderAddrs", "localhost:2001")
)

func main() {

	store := NewStore(context.Background())
	svc := NewStockService(store)

	// gRPC server here -> get request order from order service
	lis, err := net.Listen("tcp", OrderAddrs)
	if err != nil {
		log.Fatalf("Failed to listen on TCP Stock Service : %v", err.Error())
	}

	defer func() {
		if err := lis.Close(); err != nil {
			log.Fatalf("Failed to close grpc connection in Stock Service :%v", err.Error())
		}
	}()

	grpcServer := grpc.NewServer()
	NewGRPCHandler(grpcServer, svc)


	log.Print("GRPC server start!",OrderAddrs)

	// listen
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve Stock Service: %v", err)
	}

}
