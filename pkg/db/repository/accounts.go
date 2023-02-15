package repository

type Accounts struct {
	ID       int64  `bun:"id"`
	UserID   int64  `bun:"user_id"`
	Number   string `bun:"number"`
	Balance  int64  `bun:"balance"`
	Currency string `bun:"currency"`
	Lock     bool   `bun:"Lock"`
}