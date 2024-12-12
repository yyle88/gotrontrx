package gotrongrpc

import (
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	MainNetGrpc   = "grpc.trongrid.io:50051"
	ShastaNetGrpc = "grpc.shasta.trongrid.io:50051"
	NileNetGrpc   = "grpc.nile.trongrid.io:50051"
)

func NewGrpcClient(address string) (*client.GrpcClient, error) {
	conn := client.NewGrpcClient(address)
	if err := conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, errors.WithMessage(err, "wrong to start Tron gRPC client connection")
	}
	return conn, nil
}
