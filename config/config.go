// Package config defined function to load yaml configuration file
package config

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath string

// LookupEnvOrString adds the capability to get env instead of param flag
func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

// init predefined flag for passing configuration file
func init() {
	flag.StringVar(&configPath, "config", LookupEnvOrString("CONFIG", "/data/etc/config.yaml"), "Fizzbuzz microservice config path")
}

// Config is application struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Metrics  Metrics
	Jaeger   Jaeger
}

// ServerConfig is server configuration structure
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// PostgresConfig is postgres configuration structure
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	PgDriver string
}

// Metrics is prometheus configuration structure
type Metrics struct {
	URL         string
	ServiceName string
}

// Jaeger is tracing configuration structure
type Jaeger struct {
	Host        string
	ServiceName string
}

// InitConfig loads yaml configuration file
func InitConfig() (*Config, error) {
	if configPath == "" {
		return nil, errors.New("config path is not set")
	}

	base := filepath.Base(configPath)
	ext := filepath.Ext(configPath)

	cfg := &Config{}

	viper.SetConfigName(strings.TrimSuffix(base, filepath.Ext(base)))
	viper.SetConfigType(strings.TrimLeft(ext, "."))
	viper.AddConfigPath(filepath.Dir(configPath))
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil
}
