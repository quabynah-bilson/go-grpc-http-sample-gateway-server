package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// ServiceRegistrationOption is a type alias for a function that takes a pointer to either a gRPC server or a HTTP gateway server
type ServiceRegistrationOption func(*grpc.Server, *runtime.ServeMux) error

// GatewayServer is an interface for a server (gRPC or HTTP gateway)
type GatewayServer interface {
	// WithAuthServer registers the auth service with the server instance
	WithAuthServer() ServiceRegistrationOption

	// Start starts the server instance
	Start(...ServiceRegistrationOption) error

	// Stop stops the server instance
	Stop() error
}
