package repository

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

type PaymentsPostgresRepository struct {
	db bun.DB
}

func NewPaymentsPostgresRepository(db bun.DB) *PaymentsPostgresRepository {
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
		Scan(ctx, row)
	if err != nil {
		return nil, fmt.Errorf("repo get payment: %w", err)
	}

	return row, nil
}

func (r *PaymentsPostgresRepository) Update(ctx context.Context, row *model.Payment) error {
	_, err := r.db.NewUpdate().
		Model(row).
		Column("confirm").
		OmitZero().
		WherePK().
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo update payment: %w", err)
	}

	return nil
}
