package services

import (
	"context"
	"github.com/eganow/partners/sampler/api/v1/features/auth/business_logic/app"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
)

// AuthService represents a service for authentication operations. It implements the AuthSvcServer interface.
type AuthService struct {
	uc *app.AuthUseCase
	pb.UnimplementedAuthSvcServer
}

// NewAuthService creates a new AuthService instance.
func NewAuthService(uc *app.AuthUseCase) *AuthService {
	return &AuthService{
		uc: uc,
	}
}

// Login handles the login request from the client.
func (s *AuthService) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.uc.Login(req)
}
