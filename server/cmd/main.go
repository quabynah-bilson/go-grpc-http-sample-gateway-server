package main

import (
	"context"
	"github.com/eganow/partners/sampler/api/v1/cmd/server"
	"github.com/eganow/partners/sampler/api/v1/internal"
	"log"
)

// reference: https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file
func main() {

	// Initialize dependencies
	if err := internal.InitializeDependencies(); err != nil {
		log.Fatalf("failed to initialize dependencies: %v", err)
	}

	// Close databases
	defer internal.CloseDatabases()

	// Start gRPC server in a goroutine
	go startGrpcServer()

	// Start HTTP server on main thread
	startHttpGatewayServer()
}

// startGrpcServer starts the gRPC server
func startGrpcServer() {
	// create the grpc server
	grpcServer := server.NewGrpcServer()

	// set up options for service registration(s)
	opts := []server.ServiceRegistrationOption{
		grpcServer.WithAuthServer(),
		// @todo: add more services here
	}

	// start the gRPC server
	if err := grpcServer.Start(opts...); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}

	// stop the gRPC server when the function returns
	defer func(grpcServer *server.GrpcServer) {
		_ = grpcServer.Stop()
	}(grpcServer)
}

// startHttpGatewayServer starts the http gateway server
func startHttpGatewayServer() {
	// create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create the http gateway server
	httpServer := server.NewHttpGatewayServer(ctx)

	// set up options for service registration(s)
	opts := []server.ServiceRegistrationOption{
		httpServer.WithAuthServer(),
		// @todo: add more services here
	}

	// stop the http gateway server when the function returns
	defer func(httpServer *server.HttpGatewayServer) {
		_ = httpServer.Stop()
	}(httpServer)

	// start the http gateway server
	if err := httpServer.Start(opts...); err != nil {
		log.Fatalf("failed to start http gateway server: %v", err)
	}
}
