package genre

import (
	"backend/pkg/infra/api/errors"
)

var (
	ErrInvalidID          = errors.New("genre.invalid-id", "Invalid id")
	ErrInvalidName        = errors.New("genre.invalid-name", "Invalid name")
	ErrInvalidDescription = errors.New("genre.invalid-description", "Invalid description")
	ErrGenreAlreadyExists = errors.New("genre.already-exist", "Genre already exist")
	ErrGenreNotFound      = errors.New("genre.not-found", "Genre not found")
)

type Genre struct {
	ID          int64   `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
	CreatedBy   string  `db:"created_by" json:"created_by"`
	CreatedAt   string  `db:"created_at" json:"created_at"`
	UpdatedBy   *string `db:"updated_by" json:"updated_by"`
	UpdatedAt   *string `db:"updated_at" json:"updated_at"`
}

type CreateGenre struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SearchGenreQuery struct {
	Page    int `schema:"page"`
	PerPage int `schema:"per_page"`
}

type SearchGenreResult struct {
	TotalCount int64    `json:"total_count"`
	Genres     []*Genre `json:"result"`
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
}

type UpdateGenre struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (cmd *UpdateGenre) Validate() error {
	if cmd.ID == 0 {
		return ErrInvalidID
	}

	if len(cmd.Name) == 0 {
		return ErrInvalidName
	}

	if len(cmd.Description) == 0 {
		return ErrInvalidDescription
	}

	return nil
}

func (cmd *CreateGenre) Validate() error {
	if len(cmd.Name) == 0 {
		return ErrInvalidName
	}

	if len(cmd.Description) == 0 {
		return ErrInvalidName
	}

	return nil
}
