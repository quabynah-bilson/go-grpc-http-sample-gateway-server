package main

import (
	"context"
	"github.com/eganow/partners/sampler/api/v1/cmd/grpc"
	"github.com/eganow/partners/sampler/api/v1/cmd/http"
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

	// Start gRPC server
	go grpc.StartServer(
		grpc.WithAuthServer(), // register AuthServer with gRPC server
	)

	// Start HTTP server in a goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	http.StartServer(
		http.WithAuthServer(ctx), // register AuthServer with HTTP server
	)
}
