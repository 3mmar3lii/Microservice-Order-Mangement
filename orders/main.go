package main

import (
	"context"
	"log"
	"net"

	"github.com/ammar-alaa/common"
	"google.golang.org/grpc"

	pb "github.com/ammar-alaa/common/api"
	grpcutil "github.com/ammar-alaa/common/grpc"
	postgres "github.com/ammar-alaa/common/postgres"
)

var (
	grpcAddr     = common.EnvString("GRPC_ADDR", "localhost:2000")
	stockAddrStr = common.EnvString("STOCK_ADDR", "localhost:2001")
)

func main() {
	ctx := context.Background()
	pool, err := postgres.NewPool(ctx, stockAddrStr)
	if err != nil {
		log.Fatalf("failed to connect to order db: %v", err)
	}
	conn, err := grpcutil.Dial(grpcAddr)
	if err != nil {
		log.Fatal("Error while opening the stock service connection ", err)
	}

	stockClient := pb.NewStockServiceClient(conn)

	// Store creation
	store := NewPostgresOrderStore(ctx, pool)
	// Pass store to service layer
	svc := NewService(store, stockClient)

	// 1. Initialize TCP listener on port 2000
	lis, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		log.Fatalf("Failed to listen on TCP Order Service : %v", err.Error())
	}
	grpcServer := grpc.NewServer()

	NewGRPCHandler(grpcServer, svc)

	log.Print("GRPC server start!", grpcAddr)

	// listen
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve Order Service : %v", err)
	}

	defer func() {
		lis.Close()
		conn.Close()
		pool.Close()
	}()

}
