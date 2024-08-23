package dbimpl

import (
	"backend/pkg/infra/storage/db"
	"context"
	"database/sql"
)

type sqlDB struct {
	db *sql.DB
}

func NewSQL(db *sql.DB) db.DB {
	return &sqlDB{db: db}
}

func (sq *sqlDB) Close() error {
	err := sq.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (sq *sqlDB) Get(ctx context.Context, query string, args ...interface{}) interface{} {
	return sq.db.QueryRowContext(ctx, query, args...)
}

func (sq *sqlDB) Select(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	return sq.db.QueryContext(ctx, query, args...)
}

func (sq *sqlDB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return sq.db.ExecContext(ctx, query, args...)
}
