package configs

import (
	"fmt"
	"github.com/denisenkom/go-mssqldb/azuread"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

// KeyStoreConfig represents the configuration for the KeyStore.
// It contains the sensitive information for the KeyStore.
// This data will be loaded from the .env file, from a key vault or hard coded.
type KeyStoreConfig struct {
	// server
	GrpcServerPort string
	GrpcServerHost string
	HttpServerPort string
	HttpServerHost string

	// database
	DbConnUrl         string
	DbDriver          string
	dbUser            string
	dbPassword        string
	dbHost            string
	dbPort            string
	dbName            string
	DbConnMaxIdleTime time.Duration
	DbConnMaxLifetime time.Duration
	DbMaxOpenConns    int
	DbMaxIdleConns    int
	dbSslMode         bool
}

func NewKeyStoreConfig() *KeyStoreConfig {
	// load env vars
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Fatalf("failed to load env vars: %v", err)
	}

	cfg := &KeyStoreConfig{
		GrpcServerPort: "50051",
		GrpcServerHost: "0.0.0.0",
		HttpServerPort: "9900",
		HttpServerHost: "0.0.0.0",
		DbDriver:       os.Getenv("DB_DRIVER"),
		dbUser:         os.Getenv("DB_USER"),
		dbPassword:     os.Getenv("DB_PASSWORD"),
		dbHost:         os.Getenv("DB_HOST"),
		dbPort:         os.Getenv("DB_PORT"),
		dbName:         os.Getenv("DB_NAME"),
		DbConnUrl:      os.Getenv("DB_CONN_URL"),
	}

	if idleTime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_IDLE_TIME")); err == nil {
		cfg.DbConnMaxIdleTime = time.Second * time.Duration(idleTime)
	}
	if lifetime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME")); err == nil {
		cfg.DbConnMaxLifetime = time.Second * time.Duration(lifetime)
	}

	if maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS")); err == nil {
		cfg.DbMaxOpenConns = maxOpenConns
	}

	if maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS")); err == nil {
		cfg.DbMaxIdleConns = maxIdleConns
	}

	if sslMode, err := strconv.ParseBool(os.Getenv("DB_SSL_MODE")); err == nil {
		cfg.dbSslMode = sslMode
	}

	// get database driver
	if len(cfg.DbDriver) == 0 {
		cfg.DbDriver = azuread.DriverName
	}

	// create database connection url if not provided
	if len(cfg.DbConnUrl) == 0 {
		cfg.DbConnUrl = fmt.Sprintf("%s://%s:%s@%s:%s?database=%s&connection+timeout=30&encrypt=disable&trustservercertificate=%v",
			cfg.DbDriver, cfg.dbUser, cfg.dbPassword, cfg.dbHost, cfg.dbPort, cfg.dbName, cfg.dbSslMode)
	}

	return cfg
}
