package http

import (
	"fmt"
	"github.com/eganow/partners/sampler/api/v1/configs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

// StartServer starts the HTTP server
func StartServer(opts ...ServerRegistrationOption) {
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

	// register gRPC servers
	for _, opt := range opts {
		if err := opt(grpcMux); err != nil {
			log.Fatalf("failed to register gRPC server: %v", err)
		}
	}

	// Create HTTP server that listens on port 9900
	httpMux := http.NewServeMux()
	httpMux.Handle("/", grpcMux)

	// get keystore config
	cfg := configs.NewKeyStoreConfig()

	// create a listener for the HTTP server
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.HttpServerHost, cfg.HttpServerPort))
	if err != nil {
		log.Fatalf("failed to listen on HTTP server: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("Starting HTTP server on port %s", cfg.HttpServerPort)
	if err = http.Serve(lis, httpMux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
