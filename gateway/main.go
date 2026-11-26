package main

import (
	"log"
	"net/http"

	grpcutil "github.com/ammar-alaa/common/grpc"
	"github.com/ammar-alaa/common"
	_ "github.com/joho/godotenv/autoload"
	pb "github.com/ammar-alaa/common/api"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":3000")
	orderServiceAddr = common.EnvString("ORDER_ADDR","localhost:2000")
)

func main() {

	conn, err := grpcutil.Dial(orderServiceAddr)
	if err != nil {
		log.Fatal("Error while opening the order service connection ",err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal("Error while closing the grpc connection ", err)
		}
	}()

	c := pb.NewOrderServiceClient(conn)
	
	log.Println("Dilaling orders service at ",orderServiceAddr)
	mux := http.NewServeMux()
	handler := NewHander(c)
	handler.registerRoutes(mux)

	log.Printf("Starting http server at %s", httpAddr)

	// Start the server listening
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to start the http server")
	}
}
