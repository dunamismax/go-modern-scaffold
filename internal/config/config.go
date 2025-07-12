package config

import (
	"log/slog"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config holds the application configuration.
type Config struct {
	AppEnv   string `mapstructure:"APP_ENV"`
	HTTPPort int    `mapstructure:"HTTP_PORT"`
	DBURL    string `mapstructure:"DB_URL"`
	Cache    Cache  `mapstructure:",squash"`
	Redis    Redis  `mapstructure:",squash"`
}

// Cache holds the configuration for the in-memory cache.
type Cache struct {
	NumCounters int64         `mapstructure:"CACHE_NUM_COUNTERS"`
	MaxCost     int64         `mapstructure:"CACHE_MAX_COST"`
	BufferItems int64         `mapstructure:"CACHE_BUFFER_ITEMS"`
	TTL         time.Duration `mapstructure:"CACHE_TTL"`
}

// Redis holds the configuration for the Redis client.
type Redis struct {
	URL      string `mapstructure:"REDIS_URL"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

// Load loads the configuration from a .env file and environment variables.
func Load() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set default values
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("HTTP_PORT", 3000)
	viper.SetDefault("DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	// Cache defaults
	viper.SetDefault("CACHE_NUM_COUNTERS", 1e7) // 10M
	viper.SetDefault("CACHE_MAX_COST", 1<<30)   // 1GB
	viper.SetDefault("CACHE_BUFFER_ITEMS", 64)
	viper.SetDefault("CACHE_TTL", 5*time.Minute)

	// Redis defaults
	viper.SetDefault("REDIS_URL", "localhost:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			slog.Warn("config file not found, using environment variables and defaults")
		} else {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}