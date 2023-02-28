package apiv1

import (
	"fmt"
	"time"
)

type RequestResponse struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
}

type RequestLockRequest struct {
	ID        int64  `json:"id"`
	AccountID int64  `json:"account_id"`
	Status    string `json:"status"`
}

func (r *RequestLockRequest) Validate() error {
	if r.ID == 0 {
		return fmt.Errorf("id must not be empty")
	}
	return nil
}

type RequestRequest struct {
	AccountID int64 `json:"account_id"`
}

func (r *RequestRequest) Validate() error {
	if r.AccountID == 0 {
		return fmt.Errorf("empty account id")
	}
	return nil
}
