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

	ID        int64     `bun:"id,pk,autoincrement"`
	AccountID int64     `bun:"account_id,type:integer,notnull"`
	Date      time.Time `bun:"date,nullzero,notnull,default:current_timestamp"`
	Status    string    `bun:"status,type:varchar,notnull,default:'new'"`

	Account *Account `bun:"rel:belongs-to,join:account_id=id"`
}

type RequestRepository interface {
	Create(ctx context.Context, r *Request) error
	Get(ctx context.Context, id int64) (*Request, error)
	UpdateStatus(ctx context.Context, r *Request) error
}
