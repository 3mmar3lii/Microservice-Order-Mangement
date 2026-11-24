package main

import (
	"context"
	"log"
	"net"

	"github.com/ammar-alaa/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("grpcAddr", "localhost:2000")
)

func main() {


	// Store creation
	store := NewStore(context.Background())
	// Pass store to service layer
	svc := NewService(store)

	// 1. Initialize TCP listener on port 2000
	lis, err := net.Listen("tcp", grpcAddr)
	
	defer lis.Close()
	
	if err != nil {
		log.Fatalf("Failed to listen on TCP: %v", err.Error())
	}
	grpcServer := grpc.NewServer()

	NewGRPCHandler(grpcServer, svc)

	log.Print("GRPC server start!",grpcAddr)

	// listen
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
