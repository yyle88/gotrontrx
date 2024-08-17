package gotrongrpc

import (
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	MainnetGrpc   = "grpc.trongrid.io:50051"
	ShastaNetGrpc = "grpc.shasta.trongrid.io:50051"
	NileNetGrpc   = "grpc.nile.trongrid.io:50051"
)

func NewGrpcClient(address string) (*client.GrpcClient, error) {
	conn := client.NewGrpcClient(address)
	err := conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.WithMessage(err, "can not conn to tron grpc")
	}
	return conn, nil
}
