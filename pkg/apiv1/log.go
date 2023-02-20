package apiv1

import (
	"fmt"
	"time"
)

type LogResponse struct {
	ID     int64     `json:"id"`
	UserID int64     `json:"user_id"`
	Date   time.Time `json:"date"`
	Action string    `json:"action"`
}

type LogRequest struct {
	UserID int64 `json:"user_id"`
}

func (l *LogRequest) Validate() error {
	if l.UserID == 0 {
		return fmt.Errorf("empty user id")
	}
	return nil
}
