package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Payment struct {
	bun.BaseModel `bun:"table:payments,alias:p"`

	ID        int       `bun:"id"`
	AccountID string    `bun:"account_id"`
	Date      time.Time `bun:"date"`
	Sum       int64     `bun:"sum"`
	Confirm   bool      `bun:"confirm"`

	Account *Account `bun:"rel:belongs-to,join:account_id=id"`
}

type PaymentsRepository interface {
	Create(ctx context.Context, p *Payment) error
	Get(ctx context.Context, id int64) (*Payment, error)
	Update(ctx context.Context, p *Payment) error
}
