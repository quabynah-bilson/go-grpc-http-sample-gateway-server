package pkg

import pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"

// DataSource represents the data source.
type DataSource interface {
	// GetAllAccounts returns all accounts.
	GetAllAccounts() ([]*pb.Account, error)

	// GetAccountById returns an account by id.
	GetAccountById(id string) (*pb.Account, error)

	// GetAccountByEmail returns an account by email.
	GetAccountByEmail(email string) (*pb.Account, error)

	// CreateAccount creates a new account.
	CreateAccount(account *pb.CreateAccountRequest) (*pb.Account, error)
}
