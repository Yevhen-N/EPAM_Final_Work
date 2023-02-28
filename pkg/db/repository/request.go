package repository

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

type RequestPostgresRepository struct {
	db bun.IDB
}

func NewRequestPostgresRepository(db bun.IDB) *RequestPostgresRepository {
	return &RequestPostgresRepository{db: db}
}

func (r *RequestPostgresRepository) Create(ctx context.Context, row *model.Request) error {
	_, err := r.db.NewInsert().
		Model(row).
		Returning("*").
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo create request: %w", err)
	}
	return nil
}

func (r *RequestPostgresRepository) Get(ctx context.Context, id int64) (*model.Request, error) {
	row := &model.Request{}
	err := r.db.NewSelect().
		Model(row).
		Relation("Account").
		Relation("Account.User").
		Where("r.id=?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get request: %w", err)
	}
	return row, nil
}

func (r *RequestPostgresRepository) UpdateStatus(ctx context.Context, row *model.Request) error {
	_, err := r.db.NewUpdate().
		Model(row).
		Column("status").
		OmitZero().
		WherePK().
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("repo update request: %w", err)
	}
	return nil
}
