// Package config defined function to load yaml configuration file
package config

import (
	"path/filepath"
	"strings"

	"github.com/py4mac/fizzbuzz/pkg/x/errorx"
	"github.com/spf13/viper"
)

var CfgFile string

// Config is application struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Metrics  MetricsConfig
	Jaeger   JaegerConfig
}

// ServerConfig is server configuration structure
type ServerConfig struct {
	Port string
}

// PostgresConfig is postgres configuration structure
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

// MetricsConfig is prometheus configuration structure
type MetricsConfig struct {
	Port string
}

// JaegerConfig is tracing configuration structure
type JaegerConfig struct {
	Host string
	Port string
}

// InitConfig loads yaml configuration file
func InitConfig() (*Config, error) {
	base := filepath.Base(CfgFile)
	ext := filepath.Ext(CfgFile)

	cfg := &Config{}

	viper.SetConfigName(strings.TrimSuffix(base, filepath.Ext(base)))
	viper.SetConfigType(strings.TrimLeft(ext, "."))
	viper.AddConfigPath(filepath.Dir(CfgFile))
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errorx.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errorx.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil
}
