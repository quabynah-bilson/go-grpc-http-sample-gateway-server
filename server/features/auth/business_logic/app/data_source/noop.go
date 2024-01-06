package data_source

import (
	"context"
	"database/sql"
	"errors"
	"github.com/eganow/partners/sampler/api/v1/features/auth/business_logic/app/data_source/models"
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
	rows, err := ds.db.QueryContext(ctx, "select Id, Email, Name, CreatedAt from dbo.Accounts order by Id desc")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("failed to execute query: %v", err)
		return nil, errors.New("failed to execute query")
	}

	accounts := make([]*pb.Account, 0)
	if rows.Next() {
		var dbAccount models.DbAccount
		if err = rows.Scan(&dbAccount.Id, &dbAccount.Email, &dbAccount.Name, &dbAccount.CreatedAt); err != nil {
			log.Printf("failed to execute query: %v", err)
			return nil, errors.New("failed to execute query")
		}
		accounts = append(accounts, dbAccount.ToProtoAccount())
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
	var account models.DbAccount
	err := ds.db.QueryRowContext(ctx, "select Id, Email, Name, Password from dbo.Accounts where Id = $1", id).
		Scan(&account.Id, &account.Email, &account.Name, &account.Password)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("failed to execute query: %v", err)
		return nil, errors.New("failed to execute query")
	}
	log.Printf("executed query: %+v", account)

	return account.ToProtoAccount(), nil
}

// GetAccountByEmail returns an account by email.
func (ds *NoopDataSource) GetAccountByEmail(email string) (*pb.Account, error) {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stmt, err := ds.db.PrepareContext(ctx, "exec get_account_by_email @email")
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	var account models.DbAccount
	if err = stmt.QueryRowContext(ctx, sql.Named("email", email)).
		Scan(&account.Id, &account.Email, &account.Password, &account.Name, &account.CreatedAt); err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("failed to execute query: %v", err)
		return nil, errors.New("failed to execute query")
	}
	log.Printf("executed query: %v", account)

	return account.ToProtoAccount(), nil
}

// CreateAccount creates a new account.
func (ds *NoopDataSource) CreateAccount(req *pb.CreateAccountRequest) (*pb.Account, error) {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// create prepared statement
	stmt, err := ds.db.PrepareContext(ctx, "insert into dbo.Accounts(Email, Password, Name) values (@email, @password, @name)")
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	// execute query
	if _, err = stmt.ExecContext(ctx, sql.Named("email", req.GetEmail()), sql.Named("password", req.GetPassword()), sql.Named("name", req.GetName())); err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("failed to execute query: %v", err)
		return nil, errors.New("failed to execute query")
	}

	// execute query
	var createdAccount models.DbAccount
	if err = ds.db.QueryRowContext(ctx, "select Id, Email, Name from dbo.Accounts where Email = @email", sql.Named("email", req.GetEmail())).
		Scan(&createdAccount.Id, &createdAccount.Email, &createdAccount.Name); err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("failed to execute query: %v", err)
		return nil, errors.New("failed to execute query")
	}
	log.Printf("executed query: %v", createdAccount.ToProtoAccount())

	return createdAccount.ToProtoAccount(), nil
}
