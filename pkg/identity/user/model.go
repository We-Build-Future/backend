package user

import (
	"backend/pkg/util/generator"
	"errors"

	"github.com/google/uuid"
)

type Status int

const (
	Pending Status = iota + 1
	Active
	Inactive
	Deleted
)

type User struct {
	ID         int64   `db:"id" json:"id"`
	UUID       string  `db:"uuid" json:"uuid"`
	FirstName  string  `db:"first_name" json:"first_name"`
	LastName   string  `db:"last_name" json:"last_name"`
	MiddleName *string `db:"middle_name" json:"middle_name"`
	LoginName  string  `db:"login_name" json:"login_name"`
	Password   string  `db:"password" json:"-"`
	Status     Status  `db:"status" json:"status"`
	Email      *string `db:"email" json:"email"`
	Salt       string  `db:"salt" json:"-"`
	CreatedBy  string  `db:"created_by" json:"created_by"`
	CreatedAt  string  `db:"created_at" json:"created_at"`
	UpdatedBy  *string `db:"updated_by" json:"updated_by"`
	UpdatedAt  *string `db:"updated_at" json:"updated_at"`
}

type CreateUserCommand struct {
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
	LoginName  string  `json:"login_name"`
	Email      *string `json:"email"`
	Password   string  `json:"password"`
	Status     Status  `json:"status"`
	UUID       string
	Salt       string
}

type SearchUserQuery struct {
	Page    int `schema:"page"`
	PerPage int `schema:"per_page"`
}

type SearchUserResult struct {
	TotalCount int64   `json:"total_count"`
	Users      []*User `json:"result"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}

type UpdateUserCommand struct {
	ID         int64
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
}

type UpdateStatusCommand struct {
	ID     int64
	Status Status `json:"status"`
}

type UpdatePasswordCommand struct {
	ID       int64
	Password string `json:"password"`
}

type ForgotPasswordCommand struct {
	Email string `json:"email"`
}

func (cmd *CreateUserCommand) Validate() error {
	if len(cmd.LoginName) == 0 {
		return errors.New("login name is required")
	}

	if len(cmd.Password) == 0 {
		return errors.New("password is required")
	}

	if len(cmd.FirstName) == 0 {
		return errors.New("first name is required")
	}

	if len(cmd.LastName) == 0 {
		return errors.New("last name is required")
	}

	if cmd.Status == 0 {
		cmd.Status = Pending
	}

	cmd.UUID = uuid.New().String()
	cmd.Salt = generator.GenerateUniqueString(32)

	return nil
}

func (cmd *UpdateUserCommand) Validate() error {
	if cmd.ID <= 0 {
		return errors.New("id is required")
	}

	if len(cmd.FirstName) == 0 {
		return errors.New("first name is required")
	}

	if len(cmd.LastName) == 0 {
		return errors.New("last name is required")
	}

	return nil
}

func (cmd *UpdateStatusCommand) Validate() error {
	if cmd.ID <= 0 {
		return errors.New("id is required")
	}

	err := cmd.Status.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (cmd *UpdatePasswordCommand) Validate() error {
	if cmd.ID <= 0 {
		return errors.New("id is required")
	}

	if len(cmd.Password) == 0 {
		return errors.New("password is required")
	}

	return nil
}

func (cmd *ForgotPasswordCommand) Validate() error {
	if len(cmd.Email) == 0 {
		return errors.New("email is required")
	}

	return nil
}

func (s Status) Validate() error {
	if s < Pending || s > Deleted {
		return errors.New("invalid status")
	}

	return nil
}
