package movie

import "context"

type Service interface {
	Search(ctx context.Context, query *SearchMovieQuery) (*SearchMovieResult, error)
	Create(ctx context.Context, cmd *CreateMovie) error
	Update(ctx context.Context, cmd *UpdateMovie) error
	GetByID(ctx context.Context, id int64) (*Movie, error)
}
