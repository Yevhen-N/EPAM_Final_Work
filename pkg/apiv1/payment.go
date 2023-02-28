package apiv1

import (
	"fmt"
	"time"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

type PaymentResponse struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Date      time.Time `json:"date"`
	Sum       int64     `json:"sum"`
	Status    string    `json:"status"`

	Account *AccountResponse `json:"accounts"`
}

type PaymentRequest struct {
	AccountID int64  `json:"account_id"`
	Sum       int64  `json:"sum"`
	Status    string `json:"status" enum:"prepared,sent"`
}

func (p *PaymentRequest) Validate() error {
	switch p.Status {
	case model.PaymentStatusPrepared, model.PaymentStatusSent:
	// nothing to do
	default:
		return fmt.Errorf("unsupported status: %s", p.Status)
	}
	if p.AccountID == 0 {
		return fmt.Errorf("empty account id")
	}
	if p.Sum == 0 {
		return fmt.Errorf("payment sum must not be 0")
	}

	return nil
}
