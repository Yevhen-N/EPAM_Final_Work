package model

import (
	"context"

	"github.com/uptrace/bun"
)

type Card struct {
	bun.BaseModel `bun:"table:cards,alias:c"`

	ID        int64  `bun:"id,pk,autoincrement"`
	AccountID int64  `bun:"account_id,type:integer,notnull"`
	Number    string `bun:"number,type:varchar,unique,notnull"`

	Account *Account `bun:"rel:belongs-to,join:account_id=id"`
}

type CardsRepository interface {
	Create(ctx context.Context, c *Card) error
	Get(ctx context.Context, id int64) (*Card, error)
	List(ctx context.Context, accountID int64) ([]Card, error)
	Update(ctx context.Context, c *Card) error
}
