package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Log struct {
	bun.BaseModel `bun:"table:logs,alias:l"`

	ID     int64     `bun:"id,pk,autoincrement"`
	UserID int64     `bun:"user_id"`
	Date   time.Time `bun:"date"`
	Action string    `bun:"action"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
}

type LogsRepository interface {
	Create(ctx context.Context, l *Log) error
	ListByUserID(ctx context.Context, userID int64) ([]Log, error)
}
