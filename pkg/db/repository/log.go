package repository

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

type LogPostgresRepository struct {
	db bun.IDB
}

func NewLogPostgresRepository(db bun.IDB) *LogPostgresRepository {
	return &LogPostgresRepository{db: db}
}

func (r *LogPostgresRepository) Create(ctx context.Context, row *model.Log) error {
	_, err := r.db.NewInsert().
		Model(row).
		Returning("*").
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo create log: %w", err)
	}
	return nil
}

func (r *LogPostgresRepository) ListByUserID(ctx context.Context, userID int64) ([]model.Log, error) {
	rows := []model.Log{}
	err := r.db.NewSelect().
		Model(&rows).
		Where("user_id=?", userID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get log: %w", err)
	}
	return rows, nil
}
