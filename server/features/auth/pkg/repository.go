package pkg

import pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"

// Repository defines the methods that any concrete implementation of this
// repository must satisfy.
type Repository interface {
	// Login performs login logic.
	Login(*pb.LoginRequest) (*pb.AccountInfo, error)

	// GetAccounts returns all accounts.
	GetAccounts() ([]*pb.AccountInfo, error)

	// GetAccount returns an account by id.
	GetAccount(string) (*pb.AccountInfo, error)

	// CreateAccount creates a new account.
	CreateAccount(*pb.CreateAccountRequest) (*pb.AccountInfo, error)
}
