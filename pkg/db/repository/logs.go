package repository

import "time"

type Logs struct {
	ID     int64     `bun:"id"`
	UserID int64     `bun:"user_id"`
	Date   time.Time `bun:"date"`
	Action string    `bun:"action"`
}
