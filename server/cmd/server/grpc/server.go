package grpc

import (
	"fmt"
	"github.com/eganow/partners/sampler/api/v1/cmd/server"
	"github.com/eganow/partners/sampler/api/v1/cmd/server/grpc/interceptor"
	"github.com/eganow/partners/sampler/api/v1/configs"
	"github.com/eganow/partners/sampler/api/v1/features/auth/business_logic/services"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
	"github.com/eganow/partners/sampler/api/v1/internal"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// GatewayGrpcServer is the grpc server
type GatewayGrpcServer struct {
	srv *grpc.Server
	server.GatewayServer
}

// NewGrpcServer returns a new instance of the grpc server
func NewGrpcServer() *GatewayGrpcServer {
	return &GatewayGrpcServer{}
}

// Start starts the grpc server
func (g *GatewayGrpcServer) Start(opts ...server.ServiceRegistrationOption) error {
	var err error

	// create the grpc server
	g.srv = grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.LoggingUnaryInterceptor),
		grpc.StreamInterceptor(interceptor.LoggingStreamInterceptor),
	)

	// enable reflection
	reflection.Register(g.srv)

	// register services from the options
	for _, opt := range opts {
		if err = opt(g.srv, nil); err != nil {
			return err
		}
	}

	// get keystore config
	cfg := configs.NewKeyStoreConfig()

	// set up the listener for the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GrpcServerHost, cfg.GrpcServerPort))
	if err != nil {
		log.Printf("failed to listen on gRPC server: %v", err)
		return err
	}

	// Start the gRPC server
	log.Printf("Starting gRPC server on port %s", cfg.GrpcServerPort)
	if err = g.srv.Serve(lis); err != nil {
		log.Printf("failed to start gRPC server: %v", err)
	}

	return err
}

// Stop stops the grpc server
func (g *GatewayGrpcServer) Stop() error {
	log.Println("Stopping gRPC server")
	g.srv.GracefulStop()
	return nil
}

// WithAuthServer registers the auth service with the grpc server
func (*GatewayGrpcServer) WithAuthServer() server.ServiceRegistrationOption {
	return func(srv *grpc.Server, _ *runtime.ServeMux) error {
		pb.RegisterAuthSvcServer(srv, services.NewAuthService(internal.AuthInjector.UseCase))
		return nil
	}
}
