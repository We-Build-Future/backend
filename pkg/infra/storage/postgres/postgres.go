package postgres

import (
	"backend/pkg/infra/storage/db"
	"backend/pkg/infra/storage/db/dbimpl"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	log *zap.Logger
	db.DB
}

type DB interface {
	db.DB
}

func New(connection string) (DB, error) {
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

	p.DB = dbimpl.NewSQL(sqlx.NewDb(gdb, gormPostgresDB.Dialector.Name()))

	return p, nil
}
