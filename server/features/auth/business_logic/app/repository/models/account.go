package models

import pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"

// ToAccountInfo converts a pb.Account to a pb.AccountInfo.
func ToAccountInfo(a *pb.Account) *pb.AccountInfoResponse {
	return &pb.AccountInfoResponse{
		Id:        a.GetId(),
		Name:      a.GetName(),
		Email:     a.GetEmail(),
		Password:  a.GetPassword(),
		CreatedAt: a.GetCreatedAt(),
	}
}
