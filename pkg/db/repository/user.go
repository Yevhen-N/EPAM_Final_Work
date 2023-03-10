package repository

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

type UserPostgresRepository struct {
	db bun.IDB
}

func NewUserPostgresRepository(db bun.IDB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) Delete(ctx context.Context, id int64) error {

	_, err := r.db.NewDelete().
		Model(&model.User{}).
		Where("id=?", id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("repo delete user: %w", err)
	}
	return nil
}

func (r *UserPostgresRepository) Create(ctx context.Context, row *model.User) error {
	_, err := r.db.NewInsert().
		Model(row).
		Returning("*").
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo create user: %w", err)
	}
	return nil
}

func (r *UserPostgresRepository) Get(ctx context.Context, id int64) (*model.User, error) {
	row := &model.User{}
	err := r.db.NewSelect().
		Model(row).
		Relation("Accounts").
		Relation("Accounts.Cards").
		Relation("Accounts.Payments").
		Relation("Accounts.Requests").
		Relation("Logs").
		Where("u.id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get user: %w", err)
	}
	return row, nil
}

func (r *UserPostgresRepository) Update(ctx context.Context, row *model.User) error {
	_, err := r.db.NewUpdate().
		Model(row).
		OmitZero().
		WherePK().
		Returning("*").
		Exec(ctx, row)
	if err != nil {
		return fmt.Errorf("repo update user: %w", err)
	}
	return nil
}
