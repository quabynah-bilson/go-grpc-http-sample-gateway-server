package grpc

import (
	"github.com/eganow/partners/sampler/api/v1/features/auth/business_logic/services"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
	"github.com/eganow/partners/sampler/api/v1/internal"
	"google.golang.org/grpc"
)

// ServerRegistrationOption is a type alias for a function that takes a pointer to a gRPC server
type ServerRegistrationOption func(s *grpc.Server)

// WithAuthServer registers the AuthServer with the gRPC server
func WithAuthServer() ServerRegistrationOption {
	return func(s *grpc.Server) {
		pb.RegisterAuthSvcServer(s, services.NewAuthService(internal.AuthInjector.UseCase))
	}
}
