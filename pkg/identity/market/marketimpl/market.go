package marketimpl

import (
	"backend/pkg/identity/market"
	"backend/pkg/infra/storage/db"
	"context"

	"go.uber.org/zap"
)

type service struct {
	store *store
	db    db.DB
	log   *zap.Logger
}

func NewService(db db.DB) market.Service {
	return &service{
		store: newStore(db),
		db:    db,
		log:   zap.L().Named("market.service"),
	}
}

func (s *service) GetByID(ctx context.Context, id int64) (*market.Market, error) {
	return s.store.getByID(ctx, id)
}
