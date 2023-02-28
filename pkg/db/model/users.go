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

	ID       int64  `bun:"id,autoincrement,pk"`
	FullName string `bun:"full_name,type:varchar,notnull"`
	Email    string `bun:"email,type:varchar,unique,notnull"`
	Password string `bun:"password,notnull"`
	Status   string `bun:"status,type:varchar,notnull,default:'active'"`
	Role     string `bun:"role,type:varchar,notnull,default:'user'"`

	Accounts []Account `bun:"rel:has-many,join:id=user_id"`
	Logs     []Log     `bun:"rel:has-many,join:id=user_id"`
}

type UsersRepository interface {
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, u *User) error
	Get(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, u *User) error
}
