package market

import "context"

type Service interface {
	GetByID(ctx context.Context, id int64) (*Market, error)
}
