package model

import (
	"context"

	"github.com/uptrace/bun"
)

const (
	UserStatusActive  = "active"
	UserStatusBlocked = "blocked"
	UserRoleAdmin     = "admin"
	UserRoleUser      = "user"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64  `bun:"id"`
	FullName string `bun:"full_name"`
	Email    string `bun:"email"`
	Password string `bun:"password"`
	Status   string `bun:"status"`
	Role     string `bun:"role"`

	Accounts []Account `bun:"rel:has-many,join:id=user_id"`
	Logs     []Log     `bun:"rel:has-many,join:id=user_id"`
}

type UsersRepository interface {
	Create(ctx context.Context, u *User) error
	Get(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, u *User) error
}
