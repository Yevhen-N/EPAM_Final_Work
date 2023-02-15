package repository

import "time"

type Requests struct {
	ID        int64     `bum:"id"`
	AccountID string    `bum:"account_id"`
	Date      time.Time `bum:"date"`
	Status    string    `bum:"status"`
}
