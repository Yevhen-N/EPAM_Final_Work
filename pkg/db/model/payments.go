package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

const (
	PaymentStatusPrepared = "prepared"
	PaymentStatusSent     = "sent"
)

type Payment struct {
	bun.BaseModel `bun:"table:payments,alias:p"`

	ID        int64     `bun:"id"`
	AccountID int64     `bun:"account_id"`
	Date      time.Time `bun:"date"`
	Sum       int64     `bun:"sum"`
	Status    string    `bun:"status"`

	Account *Account `bun:"rel:belongs-to,join:account_id=id"`
}

type PaymentsRepository interface {
	Create(ctx context.Context, p *Payment) error
	Get(ctx context.Context, id int64) (*Payment, error)
	UpdateStatus(ctx context.Context, p *Payment) error
}
