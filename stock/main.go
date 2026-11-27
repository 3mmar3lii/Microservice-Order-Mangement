package main

import (
	"context"
	"log"
	"net"

	"github.com/ammar-alaa/common"
	_ "github.com/joho/godotenv/autoload"
	postgres "github.com/ammar-alaa/common/postgres"
	"google.golang.org/grpc"
)

var (
	OrderAddrs   = common.EnvString("OrderAddrs", "localhost:2001")
	StockConnStr = common.EnvString("STOCK_DB_URL", "postgres://ammar:secret@localhost:5432/stock_db?sslmode=disable")
)

func main() {

	ctx := context.Background()
	pool, err := postgres.NewPool(ctx, StockConnStr)
	if err != nil {
		log.Fatalf("failed to connect to stock db: %v", err)
	}

	store := NewPostgresStore(pool)
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
		pool.Close()
	}()

	grpcServer := grpc.NewServer()
	NewGRPCHandler(grpcServer, svc)

	log.Print("GRPC server start!", OrderAddrs)

	// listen
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve Stock Service: %v", err)
	}

}
