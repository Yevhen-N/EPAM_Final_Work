package apiv1

import (
	"time"
)

type LogResponse struct {
	ID     int64     `json:"id"`
	UserID int64     `json:"user_id"`
	Date   time.Time `json:"date"`
	Action string    `json:"action"`
}
