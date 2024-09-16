package movieimpl

import (
	"backend/pkg/identity/movie"
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
		log: zap.L().Named("movie.store"),
	}
}

func (s *store) create(ctx context.Context, entity *movie.Movie) error {
	rawSQL := `
		INSERT INTO "movie"(
			title,
			description,
			poster_image,
			genre,
			duration,
			director,
			release_date,
			poster_url,
			created_by,
			created_at
		)
		VALUES (
			:title,
			:description,
			:poster_image,
			:genre,
			:duration,
			:director,
			:release_date,
			:poster_url,
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

func (s *store) update(ctx context.Context, entity *movie.Movie) error {
	rawSQL := `
		UPDATE "movie"
		SET
			title = :title,
			description = :description,
			poster_image = :poster_image,
			genre = :genre,
			duration = :duration,
			director = :director,
			release_date = :release_date,
			poster_url = :poster_url,
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

func (s *store) search(ctx context.Context, query *movie.SearchMovieQuery) (*movie.SearchMovieResult, error) {
	var (
		result = &movie.SearchMovieResult{
			Movies: make([]*movie.Movie, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
	)

	sql.WriteString(`
		SELECT
			*
		FROM "movie"
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
		whereParams = append(whereParams, query.PerPage)
		whereParams = append(whereParams, offset)
	}

	err = s.db.Select(ctx, &result.Movies, sql.String(), whereParams...)
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

func (s *store) movieTaken(ctx context.Context, id int64, title string) ([]*movie.Movie, error) {
	var movies []*movie.Movie

	rawSQL := `
		SELECT
			*
		FROM "movie"
		WHERE
			id = ? OR
			title = ?
	`

	err := s.db.Select(ctx, &movies, rawSQL, id, title)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *store) getByID(ctx context.Context, id int64) (*movie.Movie, error) {
	var result movie.Movie

	rawSQL := `
		SELECT
			id,
			title,
			description,
			poster_image,
			genre,
			duration,
			director,
			release_date,
			poster_url,
			created_by,
			created_at
		FROM "movie"
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
