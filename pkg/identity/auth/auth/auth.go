package auth

import (
	"backend/pkg/identity/auth"
	"backend/pkg/identity/user"
	"backend/pkg/util/encrypt"
	"context"

	"go.uber.org/zap"
)

type service struct {
	userSvc user.Service
	log     *zap.Logger
}

func NewService(userSvc user.Service) auth.Service {
	return &service{
		userSvc: userSvc,
		log:     zap.L().Named("auth.service"),
	}
}

func (s *service) Login(ctx context.Context, cmd *auth.LoginCommand) (*auth.LoginResult, error) {
	exist, err := s.userSvc.GetByLoginName(ctx, cmd.LoginName)
	if err != nil {
		return nil, err
	}

	if exist == nil {
		return nil, auth.ErrUserOrPasswordInvalid
	}

	valid, err := encrypt.VerifyPassword(cmd.Password, exist.Salt, exist.Password)
	if err != nil {
		return nil, auth.ErrUserOrPasswordInvalid
	}

	if !valid {
		return nil, auth.ErrUserOrPasswordInvalid
	}

	return nil, nil
}
