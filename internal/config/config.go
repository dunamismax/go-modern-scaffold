package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds the application's configuration.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig holds the server's configuration.
type ServerConfig struct {
	Port int `envconfig:"SERVER_PORT" default:"8080"`
}

// DatabaseConfig holds the database's configuration.
type DatabaseConfig struct {
	DSN string `envconfig:"DATABASE_DSN" default:"host=localhost user=user password=password dbname=app port=5432 sslmode=disable"`
}

// New loads the application's configuration from environment variables.
func New() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
