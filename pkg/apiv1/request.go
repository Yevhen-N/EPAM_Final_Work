package apiv1

import (
	"time"
)

type RequestResponse struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
}
