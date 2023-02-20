package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/repository"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

// App represents main application
type App struct {
	userPostgresRepository     *repository.UserPostgresRepository
	accountPostgresRepository  *repository.AccountPostgresRepository
	cardPostgresRepository     *repository.CardPostgresRepository
	paymentsPostgresRepository *repository.PaymentsPostgresRepository
	requestPostgresRepository  *repository.RequestPostgresRepository
	logPostgresRepository      *repository.LogPostgresRepository
	srv                        *echo.Echo
}

func New(dbConn bun.IDB) (*App, error) {
	userPostgresRepository := repository.NewUserPostgresRepository(dbConn)
	accountPostgresRepository := repository.NewAccountPostgresRepository(dbConn)
	cardPostgresRepository := repository.NewCardPostgresRepository(dbConn)
	paymentsPostgresRepository := repository.NewPaymentsPostgresRepository(dbConn)
	requestPostgresRepository := repository.NewRequestPostgresRepository(dbConn)
	logPostgresRepository := repository.NewLogPostgresRepository(dbConn)

	app := &App{
		userPostgresRepository:     userPostgresRepository,
		accountPostgresRepository:  accountPostgresRepository,
		cardPostgresRepository:     cardPostgresRepository,
		paymentsPostgresRepository: paymentsPostgresRepository,
		requestPostgresRepository:  requestPostgresRepository,
		logPostgresRepository:      logPostgresRepository,
		srv:                        echo.New(),
	}

	// TODO use Middleware
	v1 := app.srv.Group("v1")

	v1.GET("/users/:id", app.GetUserHandler)
	v1.PUT("/users/lock", app.UserLockHandler)

	v1.POST("/accounts", app.CreateAccountHandler)
	v1.GET("/users/:id/accounts", app.ListAccountHandler)

	v1.POST("/card", app.CreateCardHandler)
	v1.GET("/card/:id", app.GetCardHandler)

	v1.POST("/account/:id/payment", app.CreatePaymentHandler)
	v1.GET("/accounts/:id/payments", app.ListPaymentsHandler)

	v1.POST("/request", app.CreateRequestHandler)
	v1.PUT("/request/:id/approve", app.RequestApprovedHandler)

	return app, nil
}

// Run runs application
func (a *App) Run(ctx context.Context, listenPort string) error {
	errCh := make(chan error, 1)
	go func() {
		if err := a.srv.Start(listenPort); err != nil {
			errCh <- fmt.Errorf("echo start failed %w", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-quit:
		_ = a.srv.Shutdown(ctx)
		return fmt.Errorf("got os signal %v, shutting down application", sig)
	case err := <-errCh:
		return err
	}
}

func getIDFromPath(c echo.Context) (int64, error) {
	var id int64
	err := echo.PathParamsBinder(c).Int64("id", &id).BindError()
	if err != nil {
		return 0, err
	}
	return id, nil
}
