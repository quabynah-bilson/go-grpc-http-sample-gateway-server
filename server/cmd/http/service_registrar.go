package http

import (
	"context"
	"fmt"
	"github.com/eganow/partners/sampler/api/v1/configs"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// dialOpts is a slice of grpc.DialOption (using insecure credentials)
	dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
)

// ServerRegistrationOption is a type alias for a function that takes a pointer to a mux server
type ServerRegistrationOption func(s *runtime.ServeMux) error

// WithAuthServer registers the AuthServer Handler with the mux server
func WithAuthServer(ctx context.Context) ServerRegistrationOption {
	cfg := configs.NewKeyStoreConfig()
	baseUrl := fmt.Sprintf("%s:%s", cfg.GrpcServerHost, cfg.GrpcServerPort)
	return func(s *runtime.ServeMux) error {
		return pb.RegisterAuthSvcHandlerFromEndpoint(ctx, s, baseUrl, dialOpts)
	}
}
