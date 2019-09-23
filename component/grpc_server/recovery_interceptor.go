package grpc_server

import (
	"context"
	"fmt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//errorLogger log error if panics and return internal server grpc error
func errorLogger() grpc_recovery.RecoveryHandlerFuncContext {
	return func(ctx context.Context, p interface{}) error {
		log.Errorf(ctx, "unknown error:%s", p)
		return status.Error(codes.Internal, fmt.Sprintf("unknown error: %s", p))
	}
}

func RecoveryStreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(errorLogger()))
}

func RecoveryUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(errorLogger()))
}
