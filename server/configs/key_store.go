package configs

import (
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
	DbDriver          string
	DbUser            string
	DbPassword        string
	DbHost            string
	DbPort            string
	DbName            string
	DbConnMaxIdleTime time.Duration
	DbConnMaxLifetime time.Duration
	DbMaxOpenConns    int
	DbMaxIdleConns    int
	DbSslMode         bool
}

func NewKeyStoreConfig() KeyStoreConfig {
	// load env vars
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Fatalf("failed to load env vars: %v", err)
	}

	keyStoreConfig := KeyStoreConfig{
		GrpcServerPort: "50051",
		GrpcServerHost: "0.0.0.0",
		HttpServerPort: "9900",
		HttpServerHost: "0.0.0.0",
		DbDriver:       os.Getenv("DB_DRIVER"),
		DbUser:         os.Getenv("DB_USER"),
		DbPassword:     os.Getenv("DB_PASSWORD"),
		DbHost:         os.Getenv("DB_HOST"),
		DbPort:         os.Getenv("DB_PORT"),
		DbName:         os.Getenv("DB_NAME"),
	}

	if idleTime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_IDLE_TIME")); err == nil {
		keyStoreConfig.DbConnMaxIdleTime = time.Second * time.Duration(idleTime)
	}
	if lifetime, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME")); err == nil {
		keyStoreConfig.DbConnMaxLifetime = time.Second * time.Duration(lifetime)
	}

	if maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS")); err == nil {
		keyStoreConfig.DbMaxOpenConns = maxOpenConns
	}

	if maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS")); err == nil {
		keyStoreConfig.DbMaxIdleConns = maxIdleConns
	}

	if sslMode, err := strconv.ParseBool(os.Getenv("DB_SSL_MODE")); err == nil {
		keyStoreConfig.DbSslMode = sslMode
	}

	return keyStoreConfig
}
