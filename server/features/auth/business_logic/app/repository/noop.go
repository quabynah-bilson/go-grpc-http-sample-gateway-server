package repository

import (
	"github.com/eganow/partners/sampler/api/v1/features/auth/pkg"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
	"github.com/google/uuid"
)

type NoopAuthRepository struct {
	pkg.Repository
}

// NewNoopAuthRepository creates a new NoopAuthRepository instance.
func NewNoopAuthRepository() pkg.Repository {
	return &NoopAuthRepository{}
}

func (n *NoopAuthRepository) Login(_ *pb.LoginRequest) (*pb.LoginResponse, error) {
	// @todo -> perform login logic here

	return &pb.LoginResponse{
		Token:        uuid.NewString(),
		RefreshToken: uuid.NewString(),
	}, nil
}
