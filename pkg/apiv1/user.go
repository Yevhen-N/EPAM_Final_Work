package apiv1

type UserResponse struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Lock     bool   `json:"lock"`
	Admin    bool   `json:"admin"`

	Accounts []AccountResponse `json:"accounts"`
	Logs     []LogResponse     `json:"logs"`
}

type UserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
