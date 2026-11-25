package main

import (
	"context"
	"log"
	"net"

	"github.com/ammar-alaa/common"
	"google.golang.org/grpc"

	grpcutil "github.com/ammar-alaa/common/grpc"
	pb "github.com/ammar-alaa/common/api"
)

var (
	grpcAddr  = common.EnvString("grpcAddr", "localhost:2000")
	stockAddr = common.EnvString("STOCK_ADDR", "localhost:2001")
)

func main() {

	conn, err := grpcutil.Dial(stockAddr)
	if err != nil {

	}

	stockClient := pb.NewStockServiceClient(conn)

	// Store creation
	store := NewStore(context.Background())
	// Pass store to service layer
	svc := NewService(store,stockClient)

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
	}()

}
