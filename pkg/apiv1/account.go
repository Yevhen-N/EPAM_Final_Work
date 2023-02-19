package apiv1

type AccountResponse struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Number   string `json:"number"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
	Lock     bool   `json:"lock"`

	Cards    []CardResponse    `json:"cards"`
	Payments []PaymentResponse `json:"payments"`
	Requests []RequestResponse `json:"requests"`
}

type AccountCreateRequest struct {
	Currency string `json:"currency"`
}
