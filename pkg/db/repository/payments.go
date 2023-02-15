package repository

import "time"

type Payments struct {
	ID        int       `bun:"id"`
	AccountID string    `bun:"account_id"`
	Date      time.Time `bun:"date"`
	Sum       int64     `bun:"sum"`
	Confirm   bool      `bun:"confirm"`
}
