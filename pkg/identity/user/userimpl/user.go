package userimpl

import (
	"backend/pkg/identity/user"
	"backend/pkg/infra/storage/db"
	"context"
	"time"

	"go.uber.org/zap"
)

type service struct {
	store *store
	db    db.DB
	log   *zap.Logger
}

func NewService(db db.DB) user.Service {
	return &service{
		store: newStore(db),
		db:    db,
		log:   zap.L().Named("user.service"),
	}
}

func (s *service) Create(ctx context.Context, cmd *user.CreateUserCommand) error {
	now := time.Now().Format(time.RFC3339Nano)
	err := s.store.create(ctx, &user.User{
		// UUID:       cmd.UUID,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Middle:    cmd.Middle,
		LoginName: cmd.LoginName,
		Password:  cmd.Password,
		Status:    cmd.Status,
		Email:     cmd.Email,
		// Salt:       cmd.Salt,
		CreatedBy: "",
		CreatedAt: now,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Search(ctx context.Context, query *user.SearchUserQuery) (*user.SearchUserResult, error) {
	result, err := s.store.search(ctx, query)
	if err != nil {
		return nil, err
	}

	result.Page = query.Page
	result.PerPage = query.PerPage

	return result, nil
}
