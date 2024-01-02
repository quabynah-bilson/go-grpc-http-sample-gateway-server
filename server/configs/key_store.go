package configs

// KeyStoreConfig represents the configuration for the KeyStore.
// It contains the sensitive information for the KeyStore.
// This data will be loaded from the .env file, from a key vault or hard coded.
type KeyStoreConfig struct {
	GrpcServerPort string
	GrpcServerHost string
	HttpServerPort string
	HttpServerHost string
}

func NewKeyStoreConfig() KeyStoreConfig {
	keyStoreConfig := KeyStoreConfig{
		GrpcServerPort: "50051",
		GrpcServerHost: "0.0.0.0",
		HttpServerPort: "9900",
		HttpServerHost: "0.0.0.0",
	}
	return keyStoreConfig
}
