package genre

import "context"

type Service interface {
	Create(ctx context.Context, cmd *CreateGenre) error
	Update(ctx context.Context, cmd *UpdateGenre) error
	Search(ctx context.Context, query *SearchGenreQuery) (*SearchGenreResult, error)
	GetByID(ctx context.Context, id int64) (*Genre, error)
}
