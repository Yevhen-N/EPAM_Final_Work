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
}

type AccountRequest struct {
	UserID   int64  `json:"user_id"`
	Currency string `json:"currency" enum:"USD, UAH, EUR"`
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
