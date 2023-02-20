package app

import (
	"fmt"
	"net/http"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/apiv1"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/utils/generator"

	"github.com/labstack/echo/v4"
)

// CreateAccountHandler creates account for current user
func (a *App) CreateAccountHandler(c echo.Context) error {
	req := &apiv1.AccountRequest{}
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("bind account request: %w", err)
	}

	if err := req.Validate(); err != nil {
		return fmt.Errorf("account validate error %w", err)
	}

	row := &model.Account{
		UserID:   req.UserID,
		Number:   generator.MakeIBan(),
		Balance:  0,
		Currency: req.Currency,
	}

	err := a.accountPostgresRepository.Create(c.Request().Context(), row)
	if err != nil {
		return fmt.Errorf("create account: %w", err)
	}

	if err := c.JSON(http.StatusOK, mapAccount(row)); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

// ListAccountHandler returns accounts by id
func (a *App) ListAccountHandler(c echo.Context) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return fmt.Errorf("get id from path: %w", err)
	}

	rows, err := a.accountPostgresRepository.List(c.Request().Context(), id)
	if err != nil {
		return fmt.Errorf("list account: %w", err)
	}

	resp := make([]apiv1.AccountResponse, 0, len(rows))
	for i := range rows {
		resp = append(resp, *mapAccount(&rows[i]))
	}

	if err := c.JSON(http.StatusOK, resp); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

func mapAccount(row *model.Account) *apiv1.AccountResponse {
	res := &apiv1.AccountResponse{
		ID:       row.ID,
		UserID:   row.UserID,
		Number:   row.Number,
		Balance:  row.Balance,
		Currency: row.Currency,
		Lock:     row.Lock,
	}

	res.Cards = make([]apiv1.CardResponse, 0, len(row.Cards))
	for _, card := range row.Cards {
		res.Cards = append(res.Cards, apiv1.CardResponse{
			ID:        card.ID,
			AccountID: card.AccountID,
			Number:    card.Number,
		})
	}

	res.Payments = make([]apiv1.PaymentResponse, 0, len(row.Payments))
	for _, payment := range row.Payments {
		res.Payments = append(res.Payments, apiv1.PaymentResponse{
			ID:        payment.ID,
			AccountID: payment.AccountID,
			Date:      payment.Date,
			Sum:       payment.Sum,
			Status:    payment.Status,
		})
	}

	res.Requests = make([]apiv1.RequestResponse, 0, len(row.Requests))
	for _, request := range row.Requests {
		res.Requests = append(res.Requests, apiv1.RequestResponse{
			ID:        request.ID,
			AccountID: request.AccountID,
			Date:      request.Date,
			Status:    request.Status,
		})
	}
	return res
}
