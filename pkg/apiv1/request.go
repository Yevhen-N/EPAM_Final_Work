package apiv1

import (
	"fmt"
	"time"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
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
	if r.Status != model.RequestStatusNew {
		if r.Status != model.RequestStatusApproved {
			return fmt.Errorf("payment status must be new or approved")
		}
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
