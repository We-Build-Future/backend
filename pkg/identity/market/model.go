package market

type Market struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Currency  string `db:"currency" json:"currency"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
