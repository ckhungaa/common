package grpc_client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//contextUnaryClientInterceptor copy incoming context metadata to outgoing
func contextUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ctx = metadata.NewOutgoingContext(ctx, md)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

//contextStreamClientInterceptor copy incoming context metadata to outgoing
func contextStreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (stream grpc.ClientStream, e error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ctx = metadata.NewOutgoingContext(ctx, md)
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}
