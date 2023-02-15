package repository

type Cards struct {
	ID        int64  `bun:"id"`
	AccountID string `bun:"account_id"`
	Number    string `bun:"number"`
}
