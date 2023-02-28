package repository

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

type PaymentsPostgresRepository struct {
	db bun.IDB
}

func NewPaymentsPostgresRepository(db bun.IDB) *PaymentsPostgresRepository {
	return &PaymentsPostgresRepository{db: db}
}

func (r *PaymentsPostgresRepository) Create(ctx context.Context, row *model.Payment) error {
	_, err := r.db.NewInsert().
		Model(row).
		Returning("*").
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo create payment: %w", err)
	}
	return nil
}

func (r *PaymentsPostgresRepository) Get(ctx context.Context, id int64) (*model.Payment, error) {
	row := &model.Payment{}
	err := r.db.NewSelect().
		Model(row).
		Relation("Account").
		Where("p.id=?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get payment: %w", err)
	}
	return row, nil
}

func (r *PaymentsPostgresRepository) List(ctx context.Context, accountID int64) ([]model.Payment, error) {
	rows := []model.Payment{}
	err := r.db.NewSelect().
		Model(&rows).
		Relation("Account").
		Where("account_id=?", accountID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo list payment: %w", err)
	}
	return rows, nil
}

func (r *PaymentsPostgresRepository) UpdateStatus(ctx context.Context, row *model.Payment) error {
	_, err := r.db.NewUpdate().
		Model(row).
		Column("status").
		OmitZero().
		WherePK().
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo update payment: %w", err)
	}
	return nil
}
