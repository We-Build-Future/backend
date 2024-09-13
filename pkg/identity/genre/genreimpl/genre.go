package genreimpl

import (
	"backend/pkg/config"
	"backend/pkg/identity/genre"
	"backend/pkg/infra/storage/db"
	"context"
	"time"

	"go.uber.org/zap"
)

type service struct {
	store *store
	db    db.DB
	cfg   *config.Config
	log   *zap.Logger
}

func NewService(db db.DB, cfg *config.Config) genre.Service {
	return &service{
		store: newStore(db),
		db:    db,
		cfg:   cfg,
		log:   zap.L().Named("genre.service"),
	}
}

func (s *service) Create(ctx context.Context, cmd *genre.CreateGenre) error {
	now := time.Now().Format(time.RFC3339Nano)

	result, err := s.store.genreTaken(ctx, 0, cmd.Name)
	if err != nil {
		return err
	}

	if len(result) > 0 {
		return genre.ErrGenreAlreadyExists
	}

	err = s.store.create(ctx, &genre.Genre{
		Name:        cmd.Name,
		Description: &cmd.Description,
		CreatedBy:   "",
		CreatedAt:   now,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetByID(ctx context.Context, id int64) (*genre.Genre, error) {
	result, err := s.store.getByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, genre.ErrGenreNotFound
	}

	return result, nil
}

func (s *service) Search(ctx context.Context, query *genre.SearchGenreQuery) (*genre.SearchGenreResult, error) {
	if query.Page == 0 {
		query.Page = s.cfg.Pagination.Page
	}

	if query.PerPage == 0 {
		query.PerPage = s.cfg.Pagination.PerPage
	}

	result, err := s.store.search(ctx, query)
	if err != nil {
		return nil, err
	}

	result.Page = query.Page
	result.PerPage = query.PerPage

	return result, nil
}

func (s *service) Update(ctx context.Context, cmd *genre.UpdateGenre) error {
	now := time.Now().Format(time.RFC3339Nano)

	result, err := s.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = s.store.update(ctx, &genre.Genre{
		ID:          result.ID,
		Name:        cmd.Name,
		Description: &cmd.Description,
		UpdatedAt:   &now,
	})
	if err != nil {
		return err
	}

	return nil
}
