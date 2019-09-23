package grpc_client

import (
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

var (
	retryCodes  = []codes.Code{codes.Internal, codes.Unavailable, codes.Unimplemented}
	backOffWait = 100 * time.Millisecond
)

//retryUnaryClientInterceptor retry if error
func retryUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backOffWait)),
		grpc_retry.WithCodes(retryCodes...),
		grpc_retry.WithMax(3),
	)
}

//retryStreamClientInterceptor retry if error
func retryStreamClientInterceptor() grpc.StreamClientInterceptor {
	return grpc_retry.StreamClientInterceptor(
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backOffWait)),
		grpc_retry.WithCodes(retryCodes...),
		grpc_retry.WithMax(3),
	)
}
