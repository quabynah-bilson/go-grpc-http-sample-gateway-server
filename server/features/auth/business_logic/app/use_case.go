package app

import (
	"github.com/eganow/partners/sampler/api/v1/features/auth/pkg"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
	"github.com/eganow/partners/sampler/api/v1/features/common/utils"
)

// AuthUseCase represents a use case for authentication operations.
type AuthUseCase struct {
	repo pkg.Repository
}

// NewAuthUseCase creates a new AuthUseCase instance.
func NewAuthUseCase(repo pkg.Repository) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (a *AuthUseCase) Login(req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := utils.ValidateEmail(req.GetEmail()); err != nil {
		return nil, err
	}
	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		return nil, err
	}
	return a.repo.Login(req)
}
