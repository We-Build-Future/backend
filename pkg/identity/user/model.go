package user

type Status int

const (
	Pending Status = iota + 1
	Active
	Inactive
	Deleted
)

type User struct {
	ID        int64   `db:"id" json:"id"`
	UUID      string  `db:"uuid" json:"uuid"`
	FirstName string  `db:"first_name" json:"first_name"`
	LastName  string  `db:"last_name" json:"last_name"`
	Middle    *string `db:"middle" json:"middle"`
	LoginName string  `db:"login_name" json:"login_name"`
	Password  string  `db:"password" json:"password"`
	Status    Status  `db:"status" json:"status"`
	Email     *string `db:"email" json:"email"`
	Salt      string  `db:"salt" json:"salt"`
	CreatedBy string  `db:"created_by" json:"created_by"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedBy *string `db:"updated_by" json:"updated_by"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
}

type CreateUserCommand struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Middle    *string `json:"middle_name"`
	LoginName string  `json:"login_name"`
	Email     *string `json:"email"`
	Password  string  `json:"password"`
	Status    Status  `json:"status"`
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
