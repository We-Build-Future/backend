package movieimpl

import (
	"backend/pkg/config"
	"backend/pkg/identity/movie"
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

func NewService(db db.DB, cfg *config.Config) movie.Service {
	return &service{
		store: newStore(db),
		db:    db,
		cfg:   cfg,
		log:   zap.L().Named("movie.service"),
	}
}

func (s *service) Create(ctx context.Context, cmd *movie.CreateMovie) error {
	now := time.Now().Format(time.RFC3339Nano)

	result, err := s.store.movieTaken(ctx, 0, cmd.Title)
	if err != nil {
		return err
	}

	if len(result) > 0 {
		return movie.ErrMovieAlreadyExists
	}

	err = s.store.create(ctx, &movie.Movie{
		Title:       cmd.Title,
		Description: cmd.Description,
		PosterImage: cmd.PosterImage,
		Genre:       cmd.Genre,
		Duration:    cmd.Duration,
		Director:    cmd.Director,
		ReleaseDate: cmd.ReleaseDate,
		PosterURL:   cmd.PosterURL,
		CreatedBy:   "",
		CreatedAt:   now,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Search(ctx context.Context, query *movie.SearchMovieQuery) (*movie.SearchMovieResult, error) {
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

func (s *service) Update(ctx context.Context, cmd *movie.UpdateMovie) error {
	now := time.Now().Format(time.RFC3339Nano)

	result, err := s.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = s.store.update(ctx, &movie.Movie{
		ID:          result.ID,
		Title:       cmd.Title,
		Description: cmd.Description,
		PosterImage: cmd.PosterImage,
		Genre:       cmd.Genre,
		Duration:    cmd.Duration,
		Director:    cmd.Director,
		ReleaseDate: cmd.ReleaseDate,
		PosterURL:   cmd.PosterURL,
		UpdatedAt:   &now,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetByID(ctx context.Context, id int64) (*movie.Movie, error) {
	result, err := s.store.getByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, movie.ErrMovieNotFound
	}

	return result, nil
}
