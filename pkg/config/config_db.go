package config

import (
	"backend/pkg/util/env"
	"fmt"
)

const (
	defaultDBHost     = "localhost"
	defaultDBPort     = "5432"
	defaultDBUser     = "postgres"
	defaultDBPassword = "secret"
	defaultDBName     = "identityDB"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func (p *PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC", p.Host, p.Port, p.User, p.Password, p.DB)
}

func (cfg *Config) postgresConfig() {
	cfg.Postgres.Host = env.GetEnvAsString("POSTGRES_HOST", defaultDBHost)
	cfg.Postgres.Port = env.GetEnvAsString("POSTGRES_PORT", defaultDBPort)
	cfg.Postgres.User = env.GetEnvAsString("POSTGRES_USER", defaultDBUser)
	cfg.Postgres.Password = env.GetEnvAsString("POSTGRES_PASSWORD", defaultDBPassword)
	cfg.Postgres.DB = env.GetEnvAsString("POSTGRES_DB_NAME", defaultDBName)
}
