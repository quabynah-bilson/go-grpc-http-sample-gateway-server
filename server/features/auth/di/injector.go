package di

import (
	"context"
	"database/sql"
	"github.com/eganow/partners/sampler/api/v1/configs"
	"github.com/eganow/partners/sampler/api/v1/features/auth/business_logic/app"
	"github.com/eganow/partners/sampler/api/v1/features/auth/business_logic/app/data_source"
	"github.com/eganow/partners/sampler/api/v1/features/auth/pkg"
	"log"
	"time"
)

type AuthInjector struct {
	UseCase    *app.AuthUseCase // the use case
	Repo       pkg.Repository   // the repository
	DataSource pkg.DataSource   // the data source
	DB         *sql.DB          // the database (MS SQL Server)
}

// NewAuthInjector creates a new AuthInjector instance.
func NewAuthInjector() *AuthInjector {
	injector := &AuthInjector{}

	// connect to database
	dbChannel := make(chan *sql.DB)
	go connectToDatabase(dbChannel)
	if injector.DB = <-dbChannel; injector.DB == nil {
		log.Fatal("failed to connect to database")
	}

	// create the data source
	injector.DataSource = data_source.NewNoopDataSource(injector.DB)

	// create the repository
	injector.Repo = app.NewNoopAuthRepository(injector.DataSource)

	// create the use case
	injector.UseCase = app.NewAuthUseCase(injector.Repo)

	return injector
}

// connectToDatabase connects to the MS SQL Server database.
func connectToDatabase(dbChan chan *sql.DB) {
	cfg := configs.NewKeyStoreConfig()

	// connect to database
	db, err := sql.Open(cfg.DbDriver, cfg.DbConnUrl)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		dbChan <- nil
		return
	}

	// set connection pool settings
	db.SetConnMaxIdleTime(cfg.DbConnMaxIdleTime)
	db.SetConnMaxLifetime(cfg.DbConnMaxLifetime)
	db.SetMaxIdleConns(cfg.DbMaxIdleConns)
	db.SetMaxOpenConns(cfg.DbMaxOpenConns)

	// perform health check
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		log.Printf("failed to ping connection: %v", err)
		dbChan <- nil
		return
	}

	dbChan <- db
}
