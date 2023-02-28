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

	ID        int64     `bun:"id,pk,autoincrement"`
	AccountID int64     `bun:"account_id,type:integer,notnull"`
	Date      time.Time `bun:"date,nullzero,notnull,default:current_timestamp"`
	Sum       int64     `bun:"sum,type:integer,notnull"`
	Status    string    `bun:"status,notnull,default:'prepared'"`

	Account *Account `bun:"rel:belongs-to,join:account_id=id"`
}

type PaymentsRepository interface {
	Create(ctx context.Context, p *Payment) error
	Get(ctx context.Context, id int64) (*Payment, error)
	List(ctx context.Context, accountID int64) ([]Payment, error)
	UpdateStatus(ctx context.Context, p *Payment) error
}
