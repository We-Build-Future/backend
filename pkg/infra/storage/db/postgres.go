package db

import (
	"database/sql"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	log *zap.Logger
	db  *sql.DB
}

func New(connection string) (*sql.DB, error) {
	p := &postgresDB{
		log: zap.L().Named("postgres"),
	}

	gormPostgresDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connection,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	gdb, err := gormPostgresDB.DB()
	if err != nil {
		return nil, err
	}

	p.db = gdb

	return p.db, nil
}
