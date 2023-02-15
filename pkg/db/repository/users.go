package repository

type Users struct {
	ID       int64  `bun:"ID"`
	FullName string `bun:"full_name"`
	Email    string `bun:"email"`
	Password string `bun:"password"`
	Lock     bool   `bun:"lock"`
	Admin    bool   `bun:"admin"`
}
