package repository

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

type CardPostgresRepository struct {
	db bun.IDB
}

func NewCardPostgresRepository(db bun.IDB) *CardPostgresRepository {
	return &CardPostgresRepository{db: db}
}

func (r *CardPostgresRepository) Create(ctx context.Context, row *model.Card) error {
	_, err := r.db.NewInsert().
		Model(row).
		Returning("*").
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo create card: %w", err)
	}
	return nil
}

func (r *CardPostgresRepository) Get(ctx context.Context, id int64) (*model.Card, error) {
	row := &model.Card{}
	err := r.db.NewSelect().
		Model(row).
		Relation("Account").
		Where("c.id=?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get card: %w", err)
	}
	return row, nil
}

func (r *CardPostgresRepository) List(ctx context.Context, accountID int64) ([]model.Card, error) {
	rows := []model.Card{}
	err := r.db.NewSelect().
		Model(rows).
		Relation("Account").
		Where("account_id=?", accountID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo list card: %w", err)
	}
	return rows, nil
}

func (r *CardPostgresRepository) Update(ctx context.Context, row *model.Card) error {
	_, err := r.db.NewUpdate().
		Model(row).
		OmitZero().
		WherePK().
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo update card: %w", err)
	}
	return nil
}
