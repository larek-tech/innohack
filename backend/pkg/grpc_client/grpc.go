package grpc_client

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClient holds available methods of grpc client.
type GrpcClient interface {
	GetConn() *grpc.ClientConn
	Close() error
}

// GrpcClient is an internal wrap for grpc.ClientConn.
type grpcClient struct {
	conn *grpc.ClientConn
}

// NewGrpcClient creates new GrpcClient.
func NewGrpcClient(cfg *Config, opts ...grpc.DialOption) (GrpcClient, error) {
	var grpcOpts []grpc.DialOption
	grpcOpts = append(grpcOpts, opts...)
	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	conn, err := grpc.NewClient(addr, grpcOpts...)
	if err != nil {
		return nil, err
	}

	return &grpcClient{conn: conn}, nil
}

// MustNewGrpcClientWithInsecure creates new GrpcClient with insecure credentials option enabled.
func MustNewGrpcClientWithInsecure(cfg *Config) GrpcClient {
	conn, err := NewGrpcClient(cfg, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
}

// Close grpc.ClientConn.
func (c *grpcClient) Close() error {
	return c.conn.Close()
}

// GetConn returns grpc.ClientConn.
func (c *grpcClient) GetConn() *grpc.ClientConn {
	return c.conn
}
