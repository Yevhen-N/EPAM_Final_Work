package apiv1

import "fmt"

type CardResponse struct {
	ID        int64  `json:"id"`
	AccountID int64  `json:"account_id"`
	Number    string `json:"number"`
}

type CardRequest struct {
	AccountID int64 `json:"account_id"`
}

func (c *CardRequest) Validate() error {
	if c.AccountID == 0 {
		return fmt.Errorf("empty account id")
	}
	return nil
}
