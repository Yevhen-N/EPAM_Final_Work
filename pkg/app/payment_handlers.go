package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/apiv1"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

//CreatePaymentHandler creates payment for current account
func (a *App) CreatePaymentHandler(c echo.Context) error {
	req := &apiv1.PaymentRequest{}
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("bind payment request: %w", err)
	}

	if err := req.Validate(); err != nil {
		return fmt.Errorf("payment validate error %w", err)
	}

	account, err := a.accountPostgresRepository.Get(c.Request().Context(), req.AccountID)
	if err != nil {
		return fmt.Errorf("account is not found %w", err)
	}

	if account.Status == model.AccountStatusBlocked {
		return fmt.Errorf("account already blocked")
	}

	if account.Balance < 0 || account.Balance+req.Sum < 0 {
		return fmt.Errorf("not enough money on account")
	}

	row := &model.Payment{
		AccountID: req.AccountID,
		Sum:       req.Sum,
		Status:    req.Status,
	}

	// TODO use transaction
	err = a.paymentsPostgresRepository.Create(c.Request().Context(), row)
	if err != nil {
		return fmt.Errorf("create payment: %w", err)
	}

	if req.Status == model.PaymentStatusSent {
		account.Balance += req.Sum

		err = a.accountPostgresRepository.Update(c.Request().Context(), account)
		if err != nil {
			return fmt.Errorf("failed balans update: %w", err)
		}
		logger := &model.Log{
			UserID: account.UserID,
			Action: fmt.Sprintf("User has transaction: %d m. account: # %s. Account balance: %d m.", row.Sum, account.Number, account.Balance),
		}
		if err := a.logPostgresRepository.Create(c.Request().Context(), logger); err != nil {
			return fmt.Errorf("transaction log not created: %w", err)
		}
	}

	if err := c.JSON(http.StatusOK, mapPayment(row)); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

// ListPaymentsHandler returns payments by account id
func (a *App) ListPaymentsHandler(c echo.Context) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return fmt.Errorf("get id from path: %w", err)
	}

	rows, err := a.paymentsPostgresRepository.List(c.Request().Context(), id)
	if err != nil {
		return fmt.Errorf("list account: %w", err)
	}

	resp := make([]apiv1.PaymentResponse, 0, len(rows))
	for i := range rows {
		resp = append(resp, *mapPayment(&rows[i]))
	}

	if err := c.JSON(http.StatusOK, resp); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

func mapPayment(row *model.Payment) *apiv1.PaymentResponse {
	return &apiv1.PaymentResponse{
		ID:        row.ID,
		AccountID: row.AccountID,
		Date:      row.Date,
		Sum:       row.Sum,
		Status:    row.Status,
	}
}
