package app

import (
	"context"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/repository"

	"github.com/uptrace/bun"
)

type App struct {
	userPostgresRepository     *repository.UserPostgresRepository
	accountPostgresRepository  *repository.AccountPostgresRepository
	cardPostgresRepository     *repository.CardPostgresRepository
	paymentsPostgresRepository *repository.PaymentsPostgresRepository
	requestPostgresRepository  *repository.RequestPostgresRepository
	logPostgresRepository      *repository.LogPostgresRepository
}

func New(dbConn bun.IDB) (*App, error) {
	userPostgresRepository := repository.NewUserPostgresRepository(dbConn)
	accountPostgresRepository := repository.NewAccountPostgresRepository(dbConn)
	cardPostgresRepository := repository.NewCardPostgresRepository(dbConn)
	paymentsPostgresRepository := repository.NewPaymentsPostgresRepository(dbConn)
	requestPostgresRepository := repository.NewRequestPostgresRepository(dbConn)
	logPostgresRepository := repository.NewLogPostgresRepository(dbConn)

	return &App{
		userPostgresRepository:     userPostgresRepository,
		accountPostgresRepository:  accountPostgresRepository,
		cardPostgresRepository:     cardPostgresRepository,
		paymentsPostgresRepository: paymentsPostgresRepository,
		requestPostgresRepository:  requestPostgresRepository,
		logPostgresRepository:      logPostgresRepository,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	// TODO
	return nil
}
