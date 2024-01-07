package grpc

import (
	"fmt"
	"github.com/eganow/partners/sampler/api/v1/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// StartServer starts the gRPC server
func StartServer(opts ...ServiceRegistrationOption) {
	srv := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)
	reflection.Register(srv)

	// register gRPC servers
	for _, opt := range opts {
		opt(srv)
	}

	// get keystore config
	cfg := configs.NewKeyStoreConfig()

	// set up the listener for the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GrpcServerHost, cfg.GrpcServerPort))
	if err != nil {
		log.Fatalf("failed to listen on gRPC server: %v", err)
	}

	// Start the gRPC server
	log.Printf("Starting gRPC server on port %s", cfg.GrpcServerPort)
	if err = srv.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
