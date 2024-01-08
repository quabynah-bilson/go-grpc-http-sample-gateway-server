package models

import (
	mssql "github.com/denisenkom/go-mssqldb"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
	"time"
)

// DbAccount represents an account in the database.
type DbAccount struct {
	Id        mssql.UniqueIdentifier
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time
}

// ToProtoAccount converts a DbAccount to an pb.Account.
func (a *DbAccount) ToProtoAccount() *pb.Account {
	account := &pb.Account{
		Id:       a.Id.String(),
		Name:     a.Name,
		Email:    a.Email,
		Password: a.Password,
	}

	account.CreatedAt = &pb.Timestamp{
		Seconds: a.CreatedAt.Unix(),
		Nanos:   int32(a.CreatedAt.Nanosecond()),
	}

	return account
}
