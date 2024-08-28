package marketimpl

import (
	"backend/pkg/identity/market"
	"backend/pkg/infra/storage/db"
	"context"

	"go.uber.org/zap"
)

type store struct {
	db  db.DB
	log *zap.Logger
}

func newStore(db db.DB) *store {
	return &store{
		db:  db,
		log: zap.L().Named("market.store"),
	}
}

func (s *store) getByID(ctx context.Context, id int64) (*market.Market, error) {
	var result market.Market

	rawSQL := `
		SELECT
			*
		FROM market
		WHERE 
			id = $1
	`

	err := s.db.Get(ctx, &result, rawSQL, id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
