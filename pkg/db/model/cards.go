package model

import (
	"context"

	"github.com/uptrace/bun"
)

type Card struct {
	bun.BaseModel `bun:"table:cards,alias:c"`

	ID        int64  `bun:"id"`
	AccountID string `bun:"account_id"`
	Number    string `bun:"number"`

	Account *Account `bun:"rel:belongs-to,join:account_id=id"`
}

type CardsRepository interface {
	Create(ctx context.Context, c *Card) error
	Get(ctx context.Context, id int64) (*Card, error)
	List(ctx context.Context, accountID int64) ([]Card, error)
	Update(ctx context.Context, c *Card) error
}
