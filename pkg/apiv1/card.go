package apiv1

type CardResponse struct {
	ID        int64  `json:"id"`
	AccountID int64  `json:"account_id"`
	Number    string `json:"number"`
}
