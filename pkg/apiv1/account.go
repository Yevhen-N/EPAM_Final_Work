package apiv1

import (
	"fmt"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

type AccountResponse struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Number   string `json:"number"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
	Status   string `json:"status"`

	Cards    []CardResponse    `json:"cards"`
	Payments []PaymentResponse `json:"payments"`
	Requests []RequestResponse `json:"requests"`
	User     *UserResponse     `json:"user"`
}

type AccountUpdateRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status" enums:"active, blocked"`
}

func (a *AccountUpdateRequest) Validate() error {
	if a.ID == 0 {
		return fmt.Errorf("empty account id")
	}
	switch a.Status {
	case model.AccountStatusActive, model.AccountStatusBlocked:
		// nothing to do
	default:
		return fmt.Errorf("unsupported status: %s", a.Status)
	}
	return nil
}

type AccountRequest struct {
	UserID   int64  `json:"user_id"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency" enums:"USD,UAH,EUR"`
}

func (a *AccountRequest) Validate() error {
	if a.UserID == 0 {
		return fmt.Errorf("empty user id")
	}
	switch a.Currency {
	case model.AccountCurrencyUSD, model.AccountCurrencyUAH, model.AccountCurrencyEUR:
		// nothing to do
	default:
		return fmt.Errorf("unsupported currency: %s", a.Currency)
	}
	return nil
}
