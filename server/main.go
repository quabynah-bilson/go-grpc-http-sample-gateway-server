package main

import (
	"context"
	pb "github.com/eganow/partners/sampler/api/v1/proto_gen/eganow/api"
	"github.com/eganow/partners/sampler/api/v1/services"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

const (
	grpcPort = ":50051"
	httpPort = ":9900"
)

// reference: https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file
func main() {
	authSvc := services.NewAuthService()

	// Start the gRPC and HTTP servers
	go httpServer(authSvc)
	grpcServer(authSvc)
}

// grpcServer is a function that starts a gRPC server
func grpcServer(authServer pb.AuthSvcServer) {
	srv := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)
	reflection.Register(srv)

	// register gRPC server here
	pb.RegisterAuthSvcServer(srv, authServer)

	// set up the listener for the gRPC server
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen on gRPC server: %v", err)
	}

	// Start the gRPC server
	log.Printf("Starting gRPC server on port %s", grpcPort)
	if err = srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// httpServer is a function that starts an HTTP server
func httpServer(authServer pb.AuthSvcServer) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	jsonOpts := runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	)
	grpcMux := runtime.NewServeMux(jsonOpts)
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAuthSvcHandlerServer(ctx, grpcMux, authServer); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	// Create HTTP server that listens on port 9900
	httpMux := http.NewServeMux()
	httpMux.Handle("/", grpcMux)

	// create a listener for the HTTP server
	lis, err := net.Listen("tcp", httpPort)
	if err != nil {
		log.Fatalf("failed to listen on HTTP server: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("Starting HTTP server on port %s", httpPort)
	if err = http.Serve(lis, httpMux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
