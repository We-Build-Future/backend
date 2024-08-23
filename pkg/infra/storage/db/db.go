package db

import (
	"context"
	"database/sql"
)

type DB interface {
	Get(ctx context.Context, query string, args ...interface{}) interface{}
	Select(ctx context.Context, query string, args ...interface{}) (interface{}, error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Close() error
}
