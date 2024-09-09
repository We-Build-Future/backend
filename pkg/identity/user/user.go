package user

import "context"

type Service interface {
	Create(ctx context.Context, cmd *CreateUserCommand) error
	Search(ctx context.Context, query *SearchUserQuery) (*SearchUserResult, error)
}
