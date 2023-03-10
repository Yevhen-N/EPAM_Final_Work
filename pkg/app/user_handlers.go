package app

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/apiv1"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

// DeleteUserHandler delete users
func (a *App) DeleteUserHandler(c echo.Context) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return fmt.Errorf("delete id from path: %w", err)
	}

	err = a.userPostgresRepository.Delete(c.Request().Context(), id)
	if err != nil {
		return fmt.Errorf("delete users: %w", err)
	}

	if err := c.JSON(http.StatusOK, nil); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

// CreateUserHandler create users
func (a *App) CreateUserHandler(c echo.Context) error {
	req := &apiv1.UserRequest{}
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("bind user request: %w", err)
	}

	if err := req.Validate(); err != nil {
		return fmt.Errorf("user validate error: %w", err)
	}

	row := &model.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password))),
	}

	err := a.userPostgresRepository.Create(c.Request().Context(), row)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	if err := c.JSON(http.StatusOK, mapUser(row)); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	loger := &model.Log{
		UserID: row.ID,
		Action: fmt.Sprintf("User created, name: %s, e-mail: %s.", row.FullName, row.Email),
	}

	if err := a.logPostgresRepository.Create(c.Request().Context(), loger); err != nil {
		return fmt.Errorf("user create log not created: %w", err)
	}
	return nil
}

// GetUserHandler returns user by id
func (a *App) GetUserHandler(c echo.Context) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return fmt.Errorf("get id from path: %w", err)
	}

	user, err := a.userPostgresRepository.Get(c.Request().Context(), id)
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}
	if err := c.JSON(http.StatusOK, mapUser(user)); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

func (a *App) UserRoleHandler(c echo.Context) error {
	req := &apiv1.UserStatusRequest{}
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("bind user request: %w", err)
	}

	if err := req.Validate(); err != nil {
		return fmt.Errorf("user validate error %w", err)
	}

	row := &model.User{
		ID:     req.ID,
		Status: req.Status,
	}

	err := a.userPostgresRepository.Update(c.Request().Context(), row)
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}

	if err := c.JSON(http.StatusOK, mapUser(row)); err != nil {
		return fmt.Errorf("write json response: %w", err)
	}
	return nil
}

func mapUser(row *model.User) *apiv1.UserResponse {
	res := &apiv1.UserResponse{
		ID:       row.ID,
		FullName: row.FullName,
		Email:    row.Email,
		Status:   row.Status,
		Role:     row.Role,
	}

	res.Accounts = make([]apiv1.AccountResponse, 0, len(row.Accounts))
	for _, account := range row.Accounts {
		res.Accounts = append(res.Accounts, *mapAccount(&account))
	}

	res.Logs = make([]apiv1.LogResponse, 0, len(row.Logs))
	for _, log := range row.Logs {
		res.Logs = append(res.Logs, apiv1.LogResponse{
			ID:     log.ID,
			UserID: log.UserID,
			Date:   log.Date,
			Action: log.Action,
		})
	}
	return res
}
