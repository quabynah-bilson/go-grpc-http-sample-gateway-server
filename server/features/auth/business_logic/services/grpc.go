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
	return &AuthService{uc: uc}
}

// Login handles the login request from the client.
func (s *AuthService) Login(_ context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	return s.uc.Login(req)
}

// CreateAccount handles the create account request from the client.
func (s *AuthService) CreateAccount(_ context.Context, req *pb.CreateAccountRequest) (*pb.AuthResponse, error) {
	return s.uc.CreateNewAccount(req)
}

func (s *AuthService) GetAllAccounts(_ context.Context, _ *pb.Empty) (*pb.GetAllAccountsResponse, error) {
	accounts, err := s.uc.GetAccounts()
	if err != nil {
		return nil, err
	}

	return &pb.GetAllAccountsResponse{Accounts: accounts}, nil
}
