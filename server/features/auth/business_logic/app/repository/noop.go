package repository

import (
	"github.com/eganow/partners/sampler/api/v1/features/auth/pkg"
	"github.com/eganow/partners/sampler/api/v1/features/auth/utils"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NoopAuthRepository is a repository that does nothing.
type NoopAuthRepository struct {
	ds pkg.DataSource
	pkg.Repository
}

// NewNoopAuthRepository creates a new NoopAuthRepository instance.
func NewNoopAuthRepository(ds pkg.DataSource) pkg.Repository {
	return &NoopAuthRepository{ds: ds}
}

func (n *NoopAuthRepository) Login(req *pb.LoginRequest) (*pb.Account, error) {
	account, err := n.ds.GetAccountByEmail(req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get account: %v", err)
	}

	// compare passwords
	if err = utils.ComparePasswords(req.GetPassword(), account.GetPassword()); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	return account, nil
}

func (n *NoopAuthRepository) CreateAccount(req *pb.CreateAccountRequest) (*pb.Account, error) {
	// encrypt password
	if hashedPassword, err := utils.EncryptPassword(req.GetPassword()); err != nil {
		return nil, err
	} else {
		req.Password = *hashedPassword
	}

	account, err := n.ds.CreateAccount(req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to create account: %v", err)
	}

	return account, nil
}

func (n *NoopAuthRepository) GetAccounts() ([]*pb.Account, error) {
	accounts, err := n.ds.GetAllAccounts()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get accounts: %v", err)
	}

	return accounts, nil
}

func (n *NoopAuthRepository) GetAccount(id string) (*pb.Account, error) {
	account, err := n.ds.GetAccountById(id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get account: %v", err)
	}

	return account, nil
}
