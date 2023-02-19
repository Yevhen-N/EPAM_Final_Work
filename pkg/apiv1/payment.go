package apiv1

import (
	"time"
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
	Sum    int64  `json:"sum"`
	Status string `json:"status" enum:"prepared, sent"`
}
