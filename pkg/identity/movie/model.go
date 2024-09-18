package movie

import "backend/pkg/infra/api/errors"

var (
	ErrInvalidID          = errors.New("movie.invalid-id", "Invalid id")
	ErrInvalidTitle       = errors.New("movie.invalid-title", "Invalid title")
	ErrInvalidDescription = errors.New("movie.invalid-description", "Invalid description")
	ErrMovieAlreadyExists = errors.New("movie.already-exist", "Movie already exist")
	ErrMovieNotFound      = errors.New("movie.not-found", "Movie not found")
	ErrInvalidGenre       = errors.New("movie.invalid-genre", "Invalid genre")
	ErrInvalidDuration    = errors.New("movie.invalid-duration", "Invalid duration")
	ErrInvalidDirector    = errors.New("movie.invalid-director", "Invalid director")
	ErrInvalidReleaseDate = errors.New("movie.invalid-release-date", "Invalid release date")
	ErrInvalidPosterURL   = errors.New("movie.invalid-poster-url", "Invalid poster url")
	ErrInvalidCreatedAt   = errors.New("movie.invalid-created-at", "Invalid created at")
	ErrInvalidCreatedBy   = errors.New("movie.invalid-created-by", "Invalid created by")
	ErrInvalidUpdatedAt   = errors.New("movie.invalid-updated-at", "Invalid updated at")
	ErrInvalidUpdatedBy   = errors.New("movie.invalid-updated-by", "Invalid updated by")
	ErrInvalidPosterImage = errors.New("movie.invalid-poster-image", "Invalid poster image")
)

type Genre string

const (
	ActionGenre         Genre = "Action"
	ComedyGenre         Genre = "Comedy"
	DramaGenre          Genre = "Drama"
	ScienceFictionGenre Genre = "Science Fiction"
)

type Movie struct {
	ID          int64   `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	PosterImage string  `db:"poster_image" json:"poster_image"`
	Genre       Genre   `db:"genre" json:"genre"`
	Duration    string  `db:"duration" json:"duration"`
	Director    string  `db:"director" json:"director"`
	ReleaseDate string  `db:"release_date" json:"release_date"`
	PosterURL   string  `db:"poster_url" json:"poster_url"`
	CreatedAt   string  `db:"created_at" json:"created_at"`
	CreatedBy   string  `db:"created_by" json:"created_by"`
	UpdatedAt   *string `db:"updated_at" json:"updated_at"`
	UpdatedBy   *string `db:"updated_by" json:"updated_by"`
}

type CreateMovie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PosterImage string `json:"poster_image"`
	Genre       Genre  `json:"genre"`
	Duration    string `json:"duration"`
	Director    string `json:"director"`
	ReleaseDate string `json:"release_date"`
	PosterURL   string `json:"poster_url"`
}

type SearchMovieQuery struct {
	Page    int `schema:"page"`
	PerPage int `schema:"per_page"`
}

type SearchMovieResult struct {
	TotalCount int64    `json:"total_count"`
	Movies     []*Movie `json:"result"`
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
}

type UpdateMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PosterImage string `json:"poster_image"`
	Genre       Genre  `json:"genre"`
	Duration    string `json:"duration"`
	Director    string `json:"director"`
	ReleaseDate string `json:"release_date"`
	PosterURL   string `json:"poster_url"`
}

func (cmd *UpdateMovie) Validate() error {
	if cmd.ID == 0 {
		return ErrInvalidID
	}

	if len(cmd.Title) == 0 {
		return ErrInvalidTitle
	}

	if len(cmd.Description) == 0 {
		return ErrInvalidDescription
	}

	return nil
}

func (cmd *CreateMovie) Validate() error {
	if len(cmd.Title) == 0 {
		return ErrInvalidTitle
	}

	if len(cmd.Description) == 0 {
		return ErrInvalidDescription
	}

	if cmd.Genre == "" {
		return ErrInvalidGenre
	}

	if len(cmd.Duration) == 0 {
		return ErrInvalidDuration
	}

	if len(cmd.Director) == 0 {
		return ErrInvalidDirector
	}

	if len(cmd.ReleaseDate) == 0 {
		return ErrInvalidReleaseDate
	}

	if len(cmd.PosterURL) == 0 {
		return ErrInvalidPosterURL
	}

	if len(cmd.PosterImage) == 0 {
		return ErrInvalidPosterImage
	}

	return nil
}
