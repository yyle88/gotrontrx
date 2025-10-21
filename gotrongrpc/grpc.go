// Package gotrongrpc provides gRPC client functionality for TRON blockchain network
// Supports connections to mainnet, testnet (Shasta, Nile), and custom nodes
//
// gotrongrpc 提供 TRON 区块链网络的 gRPC 客户端功能
// 支持连接到主网、测试网（Shasta、Nile）和自定义节点
package gotrongrpc

import (
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Network gRPC endpoints for TRON blockchain
// TRON 区块链网络的 gRPC 端点
const (
	MainNetGrpc   = "grpc.trongrid.io:50051"        // TRON mainnet // TRON 主网
	ShastaNetGrpc = "grpc.shasta.trongrid.io:50051" // Shasta testnet // Shasta 测试网
	NileNetGrpc   = "grpc.nile.trongrid.io:50051"   // Nile testnet // Nile 测试网
)

// NewGrpcClient creates and starts a new gRPC client connection to TRON network
// Returns initialized client ready for blockchain operations
//
// NewGrpcClient 创建并启动到 TRON 网络的 gRPC 客户端连接
// 返回已初始化的客户端，可用于区块链操作
func NewGrpcClient(address string) (*client.GrpcClient, error) {
	conn := client.NewGrpcClient(address)
	if err := conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, errors.WithMessage(err, "wrong to start Tron gRPC client connection")
	}
	return conn, nil
}
