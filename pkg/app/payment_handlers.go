package app

import (
	"fmt"
	"net/http"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/apiv1"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"

	"github.com/labstack/echo/v4"
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

	if account.Lock {
		return fmt.Errorf("account already locked")
	}

	if account.Balance <= 0 {
		if account.Balance < req.Sum*-1 {
			return fmt.Errorf("not enough money on account")
		}
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
		row := &model.Account{
			ID:      account.ID,
			Balance: account.Balance + req.Sum,
		}
		err = a.accountPostgresRepository.Update(c.Request().Context(), row)
		if err != nil {
			return fmt.Errorf("failed balans update: %w", err)
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
