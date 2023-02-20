package model

import (
	"context"

	"github.com/uptrace/bun"
)

const (
	PaymentCurrentUSD = "usd"
	PaymentCurrentUAH = "uah"
	PaymentCurrentEUR = "eur"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts,alias:a"`

	ID       int64  `bun:"id"`
	UserID   int64  `bun:"user_id"`
	Number   string `bun:"number"`
	Balance  int64  `bun:"balance"`
	Currency string `bun:"currency"`
	Lock     bool   `bun:"lock"`

	Cards    []Card    `bun:"has-many,join:id=account_id"`
	Payments []Payment `bun:"has-many,join:id=account_id"`
	Requests []Request `bun:"has-many,join:id=account_id"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
}

type AccountsRepository interface {
	Create(ctx context.Context, a *Account) error
	Get(ctx context.Context, id int64) (*Account, error)
	List(ctx context.Context, userID int64) ([]Account, error)
	Update(ctx context.Context, a *Account) error
}
