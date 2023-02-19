package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

const (
	RequestStatusNew      = "new"
	RequestStatusApproved = "approved"
)

type Request struct {
	bun.BaseModel `bun:"table:requests,alias:r"`

	ID        int64     `bum:"id"`
	AccountID int64     `bum:"account_id"`
	Date      time.Time `bum:"date"`
	Status    string    `bum:"status"`

	Account *Account `bun:"rel:belongs-to,join:account_id=id"`
}

type RequestRepository interface {
	Create(ctx context.Context, r *Request) error
	Get(ctx context.Context, id int64) (*Request, error)
	UpdateStatus(ctx context.Context, r *Request) error
}
