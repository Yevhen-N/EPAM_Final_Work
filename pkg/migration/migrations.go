package migration

import (
	"embed"
	"github.com/uptrace/bun/migrate"
)

//go:embed *.sql
var sqlMigrations embed.FS

//A collection of migrations.
var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
