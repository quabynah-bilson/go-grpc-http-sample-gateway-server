package services

import (
	"context"
	pb "github.com/eganow/partners/sampler/api/v1/proto_gen/eganow/api"
	"github.com/eganow/partners/sampler/api/v1/utils"
	"github.com/google/uuid"
)

type AuthService struct {
	pb.UnimplementedAuthSvcServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (AuthService) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := utils.ValidateEmail(req.GetEmail()); err != nil {
		return nil, err
	}
	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token:        uuid.NewString(),
		RefreshToken: uuid.NewString(),
	}, nil
}
