package userimpl

import (
	"backend/pkg/identity/user"
	"backend/pkg/infra/storage/db"
	"bytes"
	"context"
	"strings"

	"go.uber.org/zap"
)

type store struct {
	db  db.DB
	log *zap.Logger
}

func newStore(db db.DB) *store {
	return &store{
		db:  db,
		log: zap.L().Named("user.store"),
	}
}

func (s *store) create(ctx context.Context, entity *user.User) error {
	rawSQL := `
		INSERT INTO "user"(
			uuid,
			first_name,
			last_name,
			middle_name,
			login_name,
			password,
			status,
			email,
			salt,
			created_by,
			created_at
		)
		VALUES (
			:uuid,
			:first_name,
			:last_name,
			:middle_name,
			:login_name,
			:password,
			:status,
			:email,
			:salt,
			:created_by,
			:created_at
		)
	`

	_, err := s.db.NamedExec(ctx, rawSQL, entity)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) search(ctx context.Context, query *user.SearchUserQuery) (*user.SearchUserResult, error) {
	var (
		result = &user.SearchUserResult{
			Users: make([]*user.User, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
	)

	sql.WriteString(`
		SELECT
			*
		FROM "user"
	`)

	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE " + strings.Join(whereConditions, " AND "))
	}

	sql.WriteString(" ORDER BY created_at DESC")

	count, err := s.getCount(ctx, sql, whereParams)
	if err != nil {
		return nil, err
	}

	if query.PerPage > 0 {
		offset := query.PerPage * (query.Page - 1)
		sql.WriteString(" LIMIT ? OFFSET ?")
		whereParams = append(whereParams, query.PerPage, offset)
	}

	err = s.db.Select(ctx, &result.Users, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	result.TotalCount = count

	return result, nil
}

func (s *store) getCount(ctx context.Context, sql bytes.Buffer, whereParams []interface{}) (int64, error) {
	var count int64

	err := s.db.Get(ctx, &count, "SELECT COUNT(id) AS count FROM ("+sql.String()+") AS t1", whereParams...)
	if err != nil {
		return count, err
	}

	return count, nil
}
