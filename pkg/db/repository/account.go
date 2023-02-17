package repository

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

type AccountPostgresRepository struct {
	db bun.DB
}

func NewAccountPostgresRepository(db bun.DB) *AccountPostgresRepository {
	return &AccountPostgresRepository{db: db}
}

func (r *AccountPostgresRepository) Create(ctx context.Context, row *model.Account) error {
	_, err := r.db.NewInsert().
		Model(row).
		Returning("*").
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo create account: %w", err)
	}

	return nil
}

func (r *AccountPostgresRepository) Get(ctx context.Context, id int64) (*model.Account, error) {
	row := &model.Account{}
	err := r.db.NewSelect().
		Model(row).
		Relation("User").
		Relation("Cards").
		Relation("Payments").
		Relation("Requests").
		Where("id=?", id).
		Scan(ctx, row)
	if err != nil {
		return nil, fmt.Errorf("repo get account: %w", err)
	}

	return row, nil
}

func (r *AccountPostgresRepository) List(ctx context.Context, userID int64) ([]model.Account, error) {
	rows := []model.Account{}
	err := r.db.NewSelect().
		Model(rows).
		Relation("User").
		Relation("Cards").
		Relation("Payments").
		Relation("Requests").
		Where("user_id=?", userID).
		Scan(ctx, rows)
	if err != nil {
		return nil, fmt.Errorf("repo list accounts: %w", err)
	}

	return rows, nil
}

func (r *AccountPostgresRepository) Update(ctx context.Context, row *model.Account) error {
	_, err := r.db.NewUpdate().
		Model(row).
		Column("balance").
		Column("lock").
		OmitZero().
		WherePK().
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo update account: %w", err)
	}

	return nil
}
