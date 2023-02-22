package app

import (
	"fmt"
	"net/http"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/apiv1"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"

	"github.com/labstack/echo/v4"
)

// CreateRequestHandler creates request for current account
func (a *App) CreateRequestHandler(c echo.Context) error {
	req := &apiv1.RequestRequest{}
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("bind payment request: %w", err)
	}

	if err := req.Validate(); err != nil {
		return fmt.Errorf("request validate error %w", err)
	}

	row := &model.Request{
		AccountID: req.AccountID,
		Status:    model.RequestStatusNew,
	}

	err := a.requestPostgresRepository.Create(c.Request().Context(), row)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	if err := c.JSON(http.StatusOK, mapRequest(row)); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

// RequestApprovedHandler update lock
func (a *App) RequestApprovedHandler(c echo.Context) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return fmt.Errorf("get id from path: %w", err)
	}
	req, err := a.requestPostgresRepository.Get(c.Request().Context(), id)
	if err != nil {
		return fmt.Errorf("get request: %w", err)
	}
	if req.Status == model.RequestStatusApproved {
		return fmt.Errorf("request has status approved")
	}
	req.Status = model.RequestStatusApproved

	err = a.requestPostgresRepository.UpdateStatus(c.Request().Context(), req)
	if err != nil {
		return fmt.Errorf("status update: %w", err)
	}
	account := &model.Account{
		ID:     req.AccountID,
		Status: model.AccountStatusActive,
	}
	err = a.accountPostgresRepository.Update(c.Request().Context(), account)
	if err != nil {
		return fmt.Errorf("status update: %w", err)
	}

	if err := c.JSON(http.StatusOK, mapRequest(req)); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

func mapRequest(row *model.Request) *apiv1.RequestResponse {
	res := &apiv1.RequestResponse{
		ID:        row.ID,
		AccountID: row.AccountID,
		Date:      row.Date,
		Status:    row.Status,
	}
	return res
}
