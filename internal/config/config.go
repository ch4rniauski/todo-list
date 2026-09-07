package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port     string `envconfig:"POSTGRES_PORT" default:"5432"`
	User     string `envconfig:"POSTGRES_USER" default:"postgres"`
	Password string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`
	Database string `envconfig:"POSTGRES_DB" default:"catalog"`
}

func Load() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
