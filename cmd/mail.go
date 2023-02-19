package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/app"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/migration"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

func main() {
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("env DB_URI is not set")
	}
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		log.Fatal("env LISTEN_PORT is not set")
	}
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dbURI)))
	db := bun.NewDB(sqldb, pgdialect.New())
	migrator := migrate.NewMigrator(db, migration.Migrations)

	ctx := context.Background()
	err := migrator.Init(ctx)
	if err != nil {
		log.Fatal("migration init fail")
	}

	res, err := migrator.Migrate(ctx)
	if err != nil {
		log.Fatal("migration migrator fail")
	}
	log.Printf("migratoins done %s", res)

	a, err := app.New(db)
	if err != nil {
		log.Fatalf("creat app failed %s", err)
	}

	if err := a.Run(ctx); err != nil {
		log.Fatalf("run app failed %s", err)
	}
}
