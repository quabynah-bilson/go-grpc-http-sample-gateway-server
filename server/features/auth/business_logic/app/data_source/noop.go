package data_source

import (
	"context"
	"database/sql"
	"errors"
	"github.com/eganow/partners/sampler/api/v1/features/auth/pkg"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
	"log"
)

// NoopDataSource is a data source that does nothing.
type NoopDataSource struct {
	db *sql.DB
	pkg.DataSource
}

// NewNoopDataSource returns a new NoopDataSource
func NewNoopDataSource(db *sql.DB) *NoopDataSource {
	return &NoopDataSource{db: db}
}

// GetAllAccounts returns all users
func (ds *NoopDataSource) GetAllAccounts() ([]*pb.Account, error) {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// execute query
	var accounts []*pb.Account
	err := ds.db.QueryRowContext(ctx, "SELECT * FROM dbo.Accounts ORDER BY Id DESC").Scan(&accounts)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("failed to execute query: %v", err)
		return nil, errors.New("failed to execute query")
	}
	log.Printf("executed query: %+v", accounts)

	return accounts, nil
}

// GetAccountById returns an account by id.
func (ds *NoopDataSource) GetAccountById(id string) (*pb.Account, error) {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// execute query
	var account *pb.Account
	err := ds.db.QueryRowContext(ctx, "SELECT * FROM dbo.Accounts WHERE Id = $1", id).Scan(&account)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return nil, errors.New("failed to execute query")
	}
	log.Printf("executed query: %+v", account)

	return account, nil
}

// GetAccountByEmail returns an account by email.
func (ds *NoopDataSource) GetAccountByEmail(email string) (*pb.Account, error) {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// execute query
	var account *pb.Account
	err := ds.db.QueryRowContext(ctx, "SELECT * FROM dbo.Accounts WHERE Email = $1", email).Scan(&account)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return nil, errors.New("failed to execute query")
	}
	log.Printf("executed query: %+v", account)

	return account, nil
}

// CreateAccount creates a new account.
func (ds *NoopDataSource) CreateAccount(account *pb.CreateAccountRequest) (*pb.Account, error) {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// execute query
	var createdAccount *pb.Account
	err := ds.db.QueryRowContext(ctx, "INSERT INTO dbo.Accounts (Email, Password, Name) VALUES ($1, $2, $3) RETURNING *",
		account.GetEmail(), account.GetPassword(), account.GetName()).Scan(&createdAccount)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return nil, errors.New("failed to execute query")
	}
	log.Printf("executed query: %+v", createdAccount)

	return createdAccount, nil
}
