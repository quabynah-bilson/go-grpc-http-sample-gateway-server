package app

import (
	"context"
	"database/sql"
	"errors"
	"github.com/eganow/partners/sampler/api/v1/features/auth/pkg"
	pb "github.com/eganow/partners/sampler/api/v1/features/common/proto_gen/eganow/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

var (
	// queries for the data source
	getAccountByEmailQuery = "select Id, Email, Password, Name, CreatedAt from dbo.Accounts where Email = @email"
	createAccountQuery     = "insert into dbo.Accounts(Email, Password, Name) values (@email, @password, @name)"
	getAllAccountsQuery    = "select Id, Email, Name, CreatedAt from dbo.Accounts order by Id desc"
	getAccountByIdQuery    = "select Id, Email, Name, CreatedAt from dbo.Accounts where Id = @id"

	// errors
	errAccountNotFound = status.Errorf(codes.NotFound, "account not found")
	errAccountExists   = status.Errorf(codes.AlreadyExists, "account already exists")
	errFailedToExecute = status.Errorf(codes.Internal, "failed to execute query")
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

	// create prepared statement
	stmt, err := ds.db.PrepareContext(ctx, getAllAccountsQuery)
	if err != nil {
		log.Printf("failed to prepare query: %v", err)
		return nil, errFailedToExecute
	}

	// execute query
	rows, err := stmt.QueryContext(ctx, getAllAccountsQuery)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("failed to execute query: %v", err)
		return nil, errFailedToExecute
	}

	accounts := make([]*pb.Account, 0)
	if rows.Next() {
		var dbAccount *pb.Account
		if err = rows.Scan(&dbAccount.Id, &dbAccount.Email, &dbAccount.Name, &dbAccount.CreatedAt); err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, errAccountNotFound
			default:
				log.Printf("failed to execute query: %v", err)
				return nil, errFailedToExecute
			}
		}
		accounts = append(accounts, dbAccount)
	}

	return accounts, nil
}

// GetAccountById returns an account by id.
func (ds *NoopDataSource) GetAccountById(id string) (*pb.Account, error) {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// create prepared statement
	stmt, err := ds.db.PrepareContext(ctx, getAccountByIdQuery)
	if err != nil {
		log.Printf("failed to prepare query: %v", err)
		return nil, errFailedToExecute
	}

	// execute query
	var account *pb.Account
	if err = stmt.QueryRowContext(ctx, id).
		Scan(&account.Id, &account.Email, &account.Name, &account.Password); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errAccountNotFound
		default:
			log.Printf("failed to execute query: %v", err)
			return nil, errFailedToExecute
		}
	}

	return account, nil
}

// GetAccountByEmail returns an account by email.
func (ds *NoopDataSource) GetAccountByEmail(email string) (*pb.Account, error) {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stmt, err := ds.db.PrepareContext(ctx, getAccountByEmailQuery)
	if err != nil {
		log.Printf("failed to prepare query: %v", err)
		return nil, errFailedToExecute
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	account := &pb.Account{}
	if err = stmt.QueryRowContext(ctx, sql.Named("email", email)).
		Scan(&account.Id, &account.Email, &account.Password, &account.Name, &account.CreatedAt); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errAccountNotFound
		default:
			log.Printf("failed to execute query: %v", err)
			return nil, errFailedToExecute
		}
	}

	return account, nil
}

// CreateAccount creates a new account.
func (ds *NoopDataSource) CreateAccount(req *pb.CreateAccountRequest) (*pb.Account, error) {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// create prepared statement
	stmt, err := ds.db.PrepareContext(ctx, createAccountQuery)
	if err != nil {
		log.Printf("failed to prepare query: %v", err)
		return nil, errFailedToExecute
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	// execute query
	if _, err = stmt.ExecContext(ctx, sql.Named("email", req.GetEmail()), sql.Named("password",
		req.GetPassword()), sql.Named("name", req.GetName())); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errAccountExists
		default:
			log.Printf("failed to execute query: %v", err)
			return nil, errFailedToExecute
		}
	}

	// create prepared statement for select
	stmt, err = ds.db.PrepareContext(ctx, "select Id, Email, Name from dbo.Accounts where Email = @email")
	if err != nil {
		log.Printf("failed to prepare query: %v", err)
		return nil, errFailedToExecute
	}

	// execute query
	createdAccount := &pb.Account{}
	if err = stmt.QueryRowContext(ctx, sql.Named("email", req.GetEmail())).
		Scan(&createdAccount.Id, &createdAccount.Email, &createdAccount.Name); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errAccountNotFound
		default:
			log.Printf("failed to execute query: %v", err)
			return nil, errFailedToExecute
		}
	}

	return createdAccount, nil
}
