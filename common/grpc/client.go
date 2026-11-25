package grpcutil

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Dial(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}