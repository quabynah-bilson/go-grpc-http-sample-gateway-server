package models

import pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"

// ToAccountInfo converts a pb.Account to an pb.AccountInfo.
func ToAccountInfo(a *pb.Account) *pb.AccountInfo {
	return &pb.AccountInfo{
		Id:       a.Id,
		Name:     a.Name,
		Email:    a.Email,
		Password: a.Password,
	}
}
