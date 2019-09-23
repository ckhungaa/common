package grpc_client

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

//NewGRPCConnection return gRPC connection
func NewGRPCConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(
			grpc_middleware.ChainStreamClient(
				retryStreamClientInterceptor(),
				contextStreamClientInterceptor(),
			),
		),

		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				retryUnaryClientInterceptor(),
				contextUnaryClientInterceptor(),
			),
		),
	)
}
