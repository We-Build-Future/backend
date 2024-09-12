package genreimpl

import (
	"backend/pkg/identity/genre"
	"backend/pkg/infra/storage/db"
	"bytes"
	"context"
	"database/sql"
	"errors"
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
		log: zap.L().Named("genre.store"),
	}
}

func (s *store) create(ctx context.Context, entity *genre.Genre) error {

	rawSQL := `
		INSERT INTO "genre"(
			name,
			description,
			created_by,
			created_at
		)
		VALUES (
			:name,
			:description,
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

func (s *store) search(ctx context.Context, query *genre.SearchGenreQuery) (*genre.SearchGenreResult, error) {
	var (
		result = &genre.SearchGenreResult{
			Genres: make([]*genre.Genre, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
	)

	sql.WriteString(`
		SELECT
			*
		FROM "genre"
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

	err = s.db.Select(ctx, &result.Genres, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	result.TotalCount = count

	return result, nil
}

func (s *store) genreTaken(ctx context.Context, id int64, name string) ([]*genre.Genre, error) {
	var result []*genre.Genre

	rawSQL := `
		SELECT
		 	*
		FROM "genre"
		WHERE 
			id = ? OR
			name = ?
	`

	err := s.db.Select(ctx, &result, rawSQL, id, name)
	if err != nil {
		return nil, err
	}

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

func (s *store) getByID(ctx context.Context, id int64) (*genre.Genre, error) {
	var result genre.Genre

	rawSQL := `
		SELECT
			id,
			name,
			description,
			created_by,
			created_at
		FROM "genre"
		WHERE 
			id = ?
	`

	err := s.db.Get(ctx, &result, rawSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (s *store) update(ctx context.Context, entity *genre.Genre) error {
	rawSQL := `
		UPDATE "genre"
		SET
			name = :name,
			description = :description,
			updated_by = :updated_by,
			updated_at = :updated_at
		WHERE
			id = :id
	`

	_, err := s.db.NamedExec(ctx, rawSQL, entity)
	if err != nil {
		return err
	}

	return nil
}
