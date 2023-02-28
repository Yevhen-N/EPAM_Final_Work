package model

import (
	"context"

	"github.com/uptrace/bun"
)

const (
	AccountStatusActive  = "active"
	AccountStatusBlocked = "blocked"
	AccountCurrencyUSD   = "USD"
	AccountCurrencyUAH   = "UAH"
	AccountCurrencyEUR   = "EUR"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts,alias:a"`

	ID       int64  `bun:"id,autoincrement,pk"`
	UserID   int64  `bun:"user_id,type:integer"`
	Number   string `bun:"number,type:varchar,unique,notnull"`
	Balance  int64  `bun:"balance,type:integer"`
	Currency string `bun:"currency,type:varchar,default:'UAH'"`
	Status   string `bun:"status,type:varchar,default:'active'"`

	Cards    []Card    `bun:"rel:has-many,join:id=account_id"`
	Payments []Payment `bun:"rel:has-many,join:id=account_id"`
	Requests []Request `bun:"rel:has-many,join:id=account_id"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
}

type AccountsRepository interface {
	Create(ctx context.Context, a *Account) error
	Get(ctx context.Context, id int64) (*Account, error)
	List(ctx context.Context, userID int64) ([]Account, error)
	Update(ctx context.Context, a *Account) error
	UpdateStatus(ctx context.Context, a *Account) error
}
