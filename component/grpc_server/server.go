package grpc_server

import (
	"context"
	"github.com/ckhungaa/common/utils/logs"
	"github.com/google/wire"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

var (
	WireSet = wire.NewSet(
		ProvideServer,
	)
	log = logs.NewLogger("grpc_server")
)

type Server struct {
	GRPCServer *grpc.Server
	listener   net.Listener
}

//ProvideServer grpc server provider
func ProvideServer(ctx context.Context) (*Server, error) {

	port := ":8888"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			RecoveryStreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			RecoveryUnaryServerInterceptor(),
		)),
	)
	return &Server{
		GRPCServer: grpcServer,
		listener:   lis,
	}, nil
}

// Serve accepts incoming connections on the listener lis, creating a new
// ServerTransport and service goroutine for each. The service goroutines
// read gRPC requests and then call the registered handlers to reply to them.
// Serve returns when lis.Accept fails with fatal errors.  lis will be closed when
// this method returns.
// Serve will return a non-nil error unless Stop or GracefulStop is called.
func (s *Server) Serve() error {
	return s.GRPCServer.Serve(s.listener)
}
