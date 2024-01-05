package app

import (
	"github.com/eganow/partners/sampler/api/v1/features/auth/pkg"
	au "github.com/eganow/partners/sampler/api/v1/features/auth/utils"
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

func (a *AuthUseCase) Login(req *pb.LoginRequest) (*pb.AuthResponse, error) {
	if err := utils.ValidateEmail(req.GetEmail()); err != nil {
		return nil, err
	}
	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		return nil, err
	}
	account, err := a.repo.Login(req)
	if err != nil {
		return nil, err
	}

	response := &pb.AuthResponse{
		Token:        au.GenerateAccessToken(account.GetId()),
		RefreshToken: au.GenerateRefreshToken(account.GetId()),
	}

	return response, nil
}

func (a *AuthUseCase) CreateNewAccount(req *pb.CreateAccountRequest) (*pb.AuthResponse, error) {
	if err := utils.ValidateEmail(req.GetEmail()); err != nil {
		return nil, err
	}
	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		return nil, err
	}
	if err := utils.ValidateName(req.GetName()); err != nil {
		return nil, err
	}

	account, err := a.repo.CreateAccount(req)
	if err != nil {
		return nil, err
	}

	response := &pb.AuthResponse{
		Token:        au.GenerateAccessToken(account.GetId()),
		RefreshToken: au.GenerateRefreshToken(account.GetId()),
	}

	return response, nil
}

func (a *AuthUseCase) GetAccountById(id string) (*pb.Account, error) {
	if err := utils.ValidateId(id); err != nil {
		return nil, err
	}

	return a.repo.GetAccount(id)
}

func (a *AuthUseCase) GetAccounts() ([]*pb.Account, error) {
	return a.repo.GetAccounts()
}
