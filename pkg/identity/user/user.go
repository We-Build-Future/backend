package user

import "context"

type Service interface {
	Create(ctx context.Context, cmd *CreateUserCommand) error
	Search(ctx context.Context, query *SearchUserQuery) (*SearchUserResult, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, cmd *UpdateUserCommand) error
	UpdateStatus(ctx context.Context, cmd *UpdateStatusCommand) error
	UpdatePassword(ctx context.Context, cmd *UpdatePasswordCommand) error
	ForgotPassword(ctx context.Context, cmd *ForgotPasswordCommand) error
}
