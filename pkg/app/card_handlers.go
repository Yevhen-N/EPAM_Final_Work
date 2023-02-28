package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/apiv1"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/utils/generator"
)

// CreateCardHandler creates card for current account
func (a *App) CreateCardHandler(c echo.Context) error {
	req := &apiv1.CardRequest{}
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("bind card request: %w", err)
	}

	if err := req.Validate(); err != nil {
		return fmt.Errorf("card validate error %w", err)
	}

	row := &model.Card{
		AccountID: req.AccountID,
		Number:    generator.MakeCardNumber(),
	}

	err := a.cardPostgresRepository.Create(c.Request().Context(), row)
	if err != nil {
		return fmt.Errorf("create card: %w", err)
	}

	if err := c.JSON(http.StatusOK, mapCard(row)); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}

	account, err := a.accountPostgresRepository.Get(c.Request().Context(), row.AccountID)
	if err != nil {
		return fmt.Errorf("not user ID")
	}

	loger := &model.Log{
		UserID: account.UserID,
		Action: fmt.Sprintf("User created card #: %s, id: %d. Account id: %d", row.Number, row.ID, row.AccountID),
	}

	if err := a.logPostgresRepository.Create(c.Request().Context(), loger); err != nil {
		return fmt.Errorf("account block log not created: %w", err)
	}

	return nil
}

// GetCardHandler returns card by id
func (a *App) GetCardHandler(c echo.Context) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return fmt.Errorf("get id from path: %w", err)
	}

	card, err := a.cardPostgresRepository.Get(c.Request().Context(), id)
	if err != nil {
		return fmt.Errorf("get card: %w", err)
	}

	if err := c.JSON(http.StatusOK, mapCard(card)); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

func mapCard(row *model.Card) *apiv1.CardResponse {
	res := &apiv1.CardResponse{
		ID:        row.ID,
		AccountID: row.AccountID,
		Number:    row.Number,
	}
	return res
}
