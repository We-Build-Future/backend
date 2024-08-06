package config

type Config struct {
	Postgres PostgresConfig
}

func FromEnv() (*Config, error) {
	cfg := &Config{}

	cfg.postgresConfig()

	return cfg, nil
}
